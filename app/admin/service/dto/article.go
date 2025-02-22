package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"

	"time"
)

type ArticleGetPageReq struct {
	dto.Pagination `search:"-"`
	Title          string `form:"title"  search:"type:exact;column:title;table:article" comment:"标题"`
	Author         string `form:"author"  search:"type:exact;column:author;table:article" comment:"作者"`
	Content        string `form:"content"  search:"type:exact;column:content;table:article" comment:"内容"`
	Status         string `form:"status"  search:"type:exact;column:status;table:article" comment:"状态"`
	// form:"pulblishAtBeg" 与vue对应
	PublishAtBeg time.Time `form:"publishAtBeg"  search:"type:gte;column:publish_at;table:article" comment:"发布时间Beg"`
	PublishAtEnd time.Time `form:"publishAtEnd"  search:"type:lte;column:publish_at;table:article" comment:"发布时间Beg"`
	ArticleOrder
}

type ArticleOrder struct {
	//inOrder进行绑定，不存在，存在则要求类型为string
	IdOrder        string `form:"idOrder" search:"type:order;column:id;table:article"`
	TitleOrder     string `form:"titleOrder"search:"type:order;column:title;table:article"`
	AuthorOrder    string `form:"authorOrder"  search:"type:order;column:author;table:article"`
	ContentOrder   string `form:"contentOrder"  search:"type:order;column:content;table:article"`
	StatusOrder    string `form:"statusOrder"  search:"type:order;column:status;table:article"`
	PublishAtOrder string `form:"publishAtOrder"  search:"type:order;column:publish_at;table:article"`
	CreatedAtOrder string `form:"createdAtOrder"  search:"type:order;column:created_at;table:article"`
	UpdatedAtOrder string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:article"`
	DeletedAtOrder string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:article"`
	CreateByOrder  string `form:"createByOrder"  search:"type:order;column:create_by;table:article"`
	UpdateByOrder  string `form:"updateByOrder"  search:"type:order;column:update_by;table:article"`
}

func (m *ArticleGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ArticleInsertReq struct {
	Id        int       `json:"-" comment:"编码"` // 编码
	Title     string    `json:"title" comment:"标题"`
	Author    string    `json:"author" comment:"作者"`
	Content   string    `json:"content" comment:"内容"`
	Status    string    `json:"status" comment:"状态"`
	PublishAt time.Time `json:"publishAt" comment:"发布时间"`
	common.ControlBy
}

func (s *ArticleInsertReq) Generate(model *models.Article) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.Author = s.Author
	model.Content = s.Content
	model.Status = s.Status
	model.PublishAt = s.PublishAt
}

func (s *ArticleInsertReq) GetId() interface{} {
	return s.Id
}

type ArticleUpdateReq struct {
	Id        int       `uri:"id" comment:"编码"` // 编码
	Title     string    `json:"title" comment:"标题"`
	Author    string    `json:"author" comment:"作者"`
	Content   string    `json:"content" comment:"内容"`
	Status    string    `json:"status" comment:"状态"`
	PublishAt time.Time `json:"publishAt" comment:"发布时间"`
	common.ControlBy
}

func (s *ArticleUpdateReq) Generate(model *models.Article) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.Author = s.Author
	model.Content = s.Content
	model.Status = s.Status
	model.PublishAt = s.PublishAt
}

func (s *ArticleUpdateReq) GetId() interface{} {
	return s.Id
}

// ArticleGetReq 功能获取请求参数
type ArticleGetReq struct {
	Id int `uri:"id"`
}

func (s *ArticleGetReq) GetId() interface{} {
	return s.Id
}

// ArticleDeleteReq 功能删除请求参数
type ArticleDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ArticleDeleteReq) GetId() interface{} {
	return s.Ids
}
