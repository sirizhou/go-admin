package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
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
	PositionBegOrder string    `form:"positionBegOrder"search:"type:order;column:position_beg;table:formula"`
	PositionEndOrder string    `form:"positionEndOrder"search:"type:order;column:position_end;table:formula"`
}

func (m *FormulaGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type FormulaInsertReq struct {
	Id                int    `json:"-" comment:"编码"`
	Codename          string `json:"codename" comment:代号"`
	Name              string `json:"name" gcomment:名称"`
	DataType          string `json:"dataType" comment:数据类型"`
	ProcessingFormula string `json:"processingFormula" comment:处理方法"`
	FormulaParameter  string `json:"formulaParameter" comment:相关参数"`
	PositionBeg       string `json:"positionBeg" comment:起始处理位置"`
	PositionEnd       string `json:"positionEnd" comment:截止处理位置"`
}

func (s *FormulaInsertReq) Generate(model *models.Formula) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Codename = s.Codename
	model.Name = s.Name
	model.DataType = s.DataType
	model.ProcessingFormula = s.ProcessingFormula
	model.FormulaParameter = s.FormulaParameter
	model.PositionBeg = s.PositionBeg
	model.PositionEnd = s.PositionEnd
}

func (s *FormulaInsertReq) GetId() interface{} {
	return s.Id
}

type FormulaUpdateReq struct {
	Id                int    `json:"id" comment:"编码"`
	Codename          string `json:"codename" comment:代号"`
	Name              string `json:"name" gcomment:名称"`
	DataType          string `json:"dataType" comment:数据类型"`
	ProcessingFormula string `json:"processingFormula" comment:处理方法"`
	FormulaParameter  string `json:"formulaParameter" comment:相关参数"`
	// 没弄清楚为什么用intv不行 Id用int行
	PositionBeg       string    `json:"positionBeg" comment:起始处理位置"`
	PositionEnd       string    `json:"positionEnd" comment:截止处理位置"`
}

func (s *FormulaUpdateReq) Generate(model *models.Formula) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Codename = s.Codename
	model.Name = s.Name
	model.DataType = s.DataType
	model.ProcessingFormula = s.ProcessingFormula
	model.FormulaParameter = s.FormulaParameter
	model.PositionBeg = s.PositionBeg
	model.PositionEnd = s.PositionEnd
}

func (s *FormulaUpdateReq) GetId() interface{} {
	return s.Id
}

// FormulaGetReq 功能获取请求参数
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
