package models

import (
	// "gorm.io/gorm"

	"fmt"
	"go-admin/common/models"
	"time"
)

type Article struct {
	models.Model

	Title     string    `json:"title" gorm:"type:varchar(128);comment:标题"`
	Author    string    `json:"author" gorm:"type:varchar(128);comment:作者"`
	Content   string    `json:"content" gorm:"type:varchar(255);comment:内容"`
	Status    string    `json:"status" gorm:"type:int;comment:状态"`
	PublishAt time.Time `json:"publishAt" gorm:"type:timestamp;comment:发布时间"`
	// 增加一个字段，使出现在response中 ！不能在此处直接增加，会影响写入
	//ClcData string `json:"calData" gorm:"comment:计算值"`
	models.ModelTime
	models.ControlBy
}

type ArticleResponse struct {
	Article
	ClcData float64 `json:"calData"`
}

func (Article) TableName() string {
	fmt.Println("调用articel----------------------------------")
	return "article"
}

func (e *Article) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Article) GetId() interface{} {
	return e.Id
}
