package models

import (
	// "gorm.io/gorm"

	"go-admin/common/models"
	"time"
)

type DataDato struct {
	models.Model

	MongoId  string    `json:"mongoId" gorm:"type:varchar(24);comment:_id in mongoDB"`
	CreateAt time.Time `json:"createAt" gorm:"type:datetime;comment:t in mongoDB"`
	Vcid     int       `json:"vcid" gorm:"type:int;comment:vcid in mongoDB"`
	Pkid     int       `json:"pkid" gorm:"type:int;comment:vcid in mongoDB"`
	Content  string    `json:"content" gorm:"type:varchar(256);comment:content in mongoDB"`
}

type DataDatoResponse struct {
	// 增加一个用于返回计算结果的字段
	DataDato
	ClcRes float64 `json:"calRes"`
}

func (DataDato) TableName() string {
	// 避免查询时表名变为data_datos
	return "data_dato"
}

/*
func (e *DataDato) Generate() models.ActiveRecord {
	o := *e
	return &o
}
*/

func (e *DataDato) GetId() interface{} {
	return e.Id
}
