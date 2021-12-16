package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type Formula struct {
	api.Api
}

// GetPage 获取公式列表
// @Summary 获取公式列表
// @Description 获取公式列表
// @Tags 公式
// @Param title query string false "标题"
// @Param author query string false "作者"
// @Param content query string false "内容"
// @Param status query string false "状态"
// @Param publishAt query time.Time false "发布时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Formula}} "{"code": 200, "data": [...]}"
// @Router /api/v1/formula [get]
// @Security Bearer
func (e Formula) GetPage(c *gin.Context) {
	req := dto.FormulaGetPageReq{}
	s := service.Formula{}
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
	list := make([]models.Formula, 0)
	var count int64
	// 在service中实现，讲查询结果放入list
	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取公式 失败，\r\n失败信息 %s", err.Error()))
		return
	}

	//list记录response内容，对应前端response.date.list
	/*
		---------------------进行处理----------------------
			返回值放入list
	*/
	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取公式
// @Summary 获取公式
// @Description 获取公式
// @Tags 公式
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.Formula} "{"code": 200, "data": [...]}"
// @Router /api/v1/formula/{id} [get]
// @Security Bearer
func (e Formula) Get(c *gin.Context) {
	req := dto.FormulaGetReq{}
	s := service.Formula{}
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
	var object models.Formula

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取公式失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建公式
// @Summary 创建公式
// @Description 创建公式
// @Tags 公式
// @Accept application/json
// @Product application/json
// @Param data body dto.FormulaInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/formula [post]
// @Security Bearer

/*
func (e Formula) Insert(c *gin.Context) {
	req := dto.FormulaInsertReq{}
	s := service.Formula{}
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
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建公式  失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改公式
// @Summary 修改公式
// @Description 修改公式
// @Tags 公式
// @Accept application/json
// @Product application/json
// @Param data body dto.FormulaUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/formula/{id} [put]
// @Security Bearer
func (e Formula) Update(c *gin.Context) {
	req := dto.FormulaUpdateReq{}
	s := service.Formula{}
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
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改公式 失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}
*/
// Delete 删除公式
// @Summary 删除公式
// @Description 删除公式
// @Tags 公式
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/formula [delete]
// @Security Bearer
func (e Formula) Delete(c *gin.Context) {
	s := service.Formula{}
	req := dto.FormulaDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除公式失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
