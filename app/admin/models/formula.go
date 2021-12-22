package models

import (
	// "gorm.io/gorm"

	"go-admin/common/models"
)

type Formula struct {
	models.Model

	Codename          string `json:"codename" gorm:"type:varchar(32);comment:代号"`
	Name              string `json:"name" gorm:"type:varchar(64);comment:名称"`
	DataType          string `json:"dataType" gorm:"type:varchar(32);comment:数据类型"`
	ProcessingFormula string `json:"processingFormula" gorm:"type:varchar(32);comment:处理方法"`
	FormulaParameter  string `json:"formulaParameter" gorm:"type:varchar(64);comment:相关参数"`
	PositionBeg       string `json:"positionBeg" gorm"type:int;comment:起始处理位置"`
	PositionEnd       string `json:"positionEnd" gorm"type:int;comment:截止处理位置"`
}

func (Formula) TableName() string {
	return "formula"
}

/*
func (e *Formula) Generate() models.ActiveRecord {
	o := *e
	return &o
}
*/

func (e *Formula) GetId() interface{} {
	return e.Id
}
