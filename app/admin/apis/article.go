package apis

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type Article struct {
	api.Api
}

// GetPage 获取文章列表
// @Summary 获取文章列表
// @Description 获取文章列表
// @Tags 文章
// @Param title query string false "标题"
// @Param author query string false "作者"
// @Param content query string false "内容"
// @Param status query string false "状态"
// @Param publishAt query time.Time false "发布时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Article}} "{"code": 200, "data": [...]}"
// @Router /api/v1/article [get]
// @Security Bearer
func (e Article) GetPage(c *gin.Context) {
	req := dto.ArticleGetPageReq{}
	s := service.Article{}
	err := e.MakeContext(c). //增加Http上下文
					Bind(&req).              //将请求绑定到查表需求
					MakeService(&s.Service). //e赋值给s
					Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 绑定数据库
	s.Service.Orm = sdk.Runtime.GetDb()["ace"]
	p := actions.GetPermissionFromContext(c)
	list := make([]models.Article, 0)
	var count int64
	// 在service中实现，讲查询结果放入list
	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取文章 失败，\r\n失败信息 %s", err.Error()))
		return
	}
	listAdd := make([]models.ArticleResponse, len(list))
	for index, value := range list {
		listAdd[index].Article = value
		cal, _ := strconv.ParseFloat(value.Content, 64)
		cal /= 4000000000.0
		listAdd[index].ClcData = cal
	}
	//list记录response内容，对应前端response.date.list
	/*
		---------------------进行处理----------------------
			返回值放入list
	*/
	e.PageOK(listAdd, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取文章
// @Summary 获取文章
// @Description 获取文章
// @Tags 文章
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.Article} "{"code": 200, "data": [...]}"
// @Router /api/v1/article/{id} [get]
// @Security Bearer
func (e Article) Get(c *gin.Context) {
	req := dto.ArticleGetReq{}
	s := service.Article{}
	err := e.MakeContext(c).
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 绑定数据库
	s.Service.Orm = sdk.Runtime.GetDb()["ace"]
	var object models.Article

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取文章失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建文章
// @Summary 创建文章
// @Description 创建文章
// @Tags 文章
// @Accept application/json
// @Product application/json
// @Param data body dto.ArticleInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/article [post]
// @Security Bearer
func (e Article) Insert(c *gin.Context) {
	req := dto.ArticleInsertReq{}
	s := service.Article{}
	err := e.MakeContext(c).
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 绑定数据库
	s.Service.Orm = sdk.Runtime.GetDb()["ace"]
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建文章  失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改文章
// @Summary 修改文章
// @Description 修改文章
// @Tags 文章
// @Accept application/json
// @Product application/json
// @Param data body dto.ArticleUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/article/{id} [put]
// @Security Bearer
func (e Article) Update(c *gin.Context) {
	req := dto.ArticleUpdateReq{}
	s := service.Article{}
	err := e.MakeContext(c).
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 绑定数据库
	s.Service.Orm = sdk.Runtime.GetDb()["ace"]
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改文章 失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
	fmt.Println("修改成功------------------------------------------------------------------")
}

// Delete 删除文章
// @Summary 删除文章
// @Description 删除文章
// @Tags 文章
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/article [delete]
// @Security Bearer
func (e Article) Delete(c *gin.Context) {
	s := service.Article{}
	req := dto.ArticleDeleteReq{}
	err := e.MakeContext(c).
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 绑定数据库
	s.Service.Orm = sdk.Runtime.GetDb()["ace"]

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除文章失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
