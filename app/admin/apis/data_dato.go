package apis

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type DataDato struct {
	api.Api
}

// GetPage 获取遥测包列表
// @Summary 获取遥测包列表
// @Description 获取遥测包列表
// @Tags 遥测包
// @Param createAt query time.Time false "创建时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.DataDato}} "{"code": 200, "data": [...]}"
// @Router /api/v1/dato [get]
// @Security Bearer
func (e DataDato) GetPage(c *gin.Context) {
	req := dto.DataDatoGetPageReq{}
	s := service.DataDato{}
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
	s.Service.Orm = sdk.Runtime.GetDb()["data"]
	p := actions.GetPermissionFromContext(c)
	list := make([]models.DataDato, 0)
	var count int64
	// 在service中实现，讲查询结果放入list
	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取遥测包 失败，\r\n失败信息 %s", err.Error()))
		return
	}
	listAdd := make([]models.DataDatoResponse, len(list))
	for index, value := range list {
		listAdd[index].DataDato = value
		cal, _ := strconv.ParseFloat(value.Content, 64)
		cal = cal * cal / 100
		listAdd[index].ClcRes = cal
	}
	//list记录response内容，对应前端response.date.list
	/*
		---------------------进行处理----------------------
			返回值放入list
	*/
	e.PageOK(listAdd, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取遥测包
// @Summary 获取遥测包
// @Description 获取遥测包
// @Tags 遥测包
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.DataDato} "{"code": 200, "data": [...]}"
// @Router /api/v1/article/{id} [get]
// @Security Bearer
func (e DataDato) Get(c *gin.Context) {
	req := dto.DataDatoGetReq{}
	s := service.DataDato{}
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
	s.Service.Orm = sdk.Runtime.GetDb()["data"]
	var object models.DataDato

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("测包失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// 不需要update和create

// Delete 删除遥测包
// @Summary 删除遥测包
// @Description 删除遥测包
// @Tags 遥测包
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/article [delete]
// @Security Bearer
func (e DataDato) Delete(c *gin.Context) {
	s := service.DataDato{}
	req := dto.DataDatoDeleteReq{}
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
	s.Service.Orm = sdk.Runtime.GetDb()["data"]

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除遥测包失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
