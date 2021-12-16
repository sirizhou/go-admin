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

type Formula struct {
	service.Service
}

// GetPage 获取Formula列表
func (e *Formula) GetPage(c *dto.FormulaGetPageReq, p *actions.DataPermission, list *[]models.Formula, count *int64) error {
	var err error
	var data models.Formula
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
		e.Log.Errorf("FormulaService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Formula对象
func (e *Formula) Get(d *dto.FormulaGetReq, p *actions.DataPermission, model *models.Formula) error {
	var data models.Formula

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetFormula error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

/*
// Insert 创建Formula对象
func (e *Formula) Insert(c *dto.FormulaInsertReq) error {
	var err error
	var data models.Formula
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("FormulaService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Formula对象
func (e *Formula) Update(c *dto.FormulaUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Formula{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		e.Log.Errorf("FormulaService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}
*/
// Remove 删除Formula
func (e *Formula) Remove(d *dto.FormulaDeleteReq, p *actions.DataPermission) error {
	var data models.Formula

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveFormula error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
