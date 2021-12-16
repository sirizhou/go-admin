package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type DataDato struct {
	service.Service
}

// GetPage 获取DataDato列表
func (e *DataDato) GetPage(c *dto.DataDatoGetPageReq, p *actions.DataPermission, list *[]models.DataDato, count *int64) error {
	var err error
	var data models.DataDato
	// Orm 是 *gorm.DB类型
	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),            //** 包含搜索条件
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()), //分页
			actions.Permission(data.TableName(), p),          //搜索权限
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("DataDatoService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取DataDato对象
func (e *DataDato) Get(d *dto.DataDatoGetReq, p *actions.DataPermission, model *models.DataDato) error {
	var data models.DataDato

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetDataDato error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Remove 删除DataDato
func (e *DataDato) Remove(d *dto.DataDatoDeleteReq, p *actions.DataPermission) error {
	var data models.DataDato

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveDataDato error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
