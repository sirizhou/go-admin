package dto

import (
	"go-admin/common/dto"
)

type FormulaGetPageReq struct {
	dto.Pagination `search:"-"`
	Id             int    `form:"id"  search:"type:exact;column:id;table:formula" comment:"编码"`
	Codename       string `form:"codename"  search:"type:contains;column:codename;table:formula" comment:"代号"`
	Name           string `form:"name"  search:"type:contains;column:name;table:formula" comment:"名称"`
	FormulaOrder
}

type FormulaOrder struct {
	//inOrder进行绑定，不存在，存在则要求类型为string
	IdOrder          string `form:"idOrder" search:"type:order;column:id;table:formula"`
	CodeNameOrder    string `form:"codenameOrder"search:"type:order;column:codename;table:formula"`
	PositionBegOrder string `form:"positionBegOrder"search:"type:order;column:position_beg;table:formula"`
	PositionEndOrder string `form:"positionEndOrder"search:"type:order;column:position_end;table:formula"`
}

func (m *FormulaGetPageReq) GetNeedSearch() interface{} {
	return *m
}

/*
type FormulaInsertReq struct {
	Id        int       `json:"-" comment:"编码"` // 编码
	Title     string    `json:"title" comment:"标题"`
	Author    string    `json:"author" comment:"作者"`
	Content   string    `json:"content" comment:"内容"`
	Status    string    `json:"status" comment:"状态"`
	PublishAt time.Time `json:"publishAt" comment:"发布时间"`
}

func (s *FormulaInsertReq) Generate(model *models.Formula) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.Author = s.Author
	model.Content = s.Content
	model.Status = s.Status
	model.PublishAt = s.PublishAt
}

func (s *FormulaInsertReq) GetId() interfdata{} {
	return s.Id
}

type FormulaUpdateReq struct {
	Id        int       `uri:"id" comment:"编码"` // 编码
	Title     string    `json:"title" comment:"标题"`
	Author    string    `json:"author" comment:"作者"`
	Content   string    `json:"content" comment:"内容"`
	Status    string    `json:"status" comment:"状态"`
	PublishAt time.Time `json:"publishAt" comment:"发布时间"`
	common.ControlBy
}

func (s *FormulaUpdateReq) Generate(model *models.Formula) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.Author = s.Author
	model.Content = s.Content
	model.Status = s.Status
	model.PublishAt = s.PublishAt
}

func (s *FormulaUpdateReq) GetId() interfdata{} {
	return s.Id
}

// FormulaGetReq 功能获取请求参数
type FormulaGetReq struct {
	Id int `uri:"id"`
}

func (s *FormulaGetReq) GetId() interfdata{} {
	return s.Id
}
*/

type FormulaGetReq struct {
	Id int `uri:"id"`
}

func (s *FormulaGetReq) GetId() interface{} {
	return s.Id
}

// FormulaDeleteReq 功能删除请求参数
type FormulaDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *FormulaDeleteReq) GetId() interface{} {
	return s.Ids
}
