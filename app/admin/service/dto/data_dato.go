package dto

import (
	"go-admin/common/dto"

	"time"
)

type DataDatoGetPageReq struct {
	dto.Pagination `search:"-"`
	Id             int       `form:"id"  search:"type:exact;column:id;table:data_dato" comment:"编码"`
	MongoId        string    `form:"mongoId"  search:"type:exact;column:mongo_id;table:data_dato" comment:"_id in mongoDB"`
	Vcid           int       `form:"vcid"  search:"type:exact;column:vcid;table:data_dato" comment:"vcid in mongoDB"`
	Pkid           int       `form:"pkid"  search:"type:exact;column:pkid;table:data_dato" comment:"pkid in mongoDB"`
	CreateAtBeg    time.Time `form:"createAtBeg"  search:"type:gte;column:create_at;table:data_dato" comment:"t in mongoDB"`
	CreateAtEnd    time.Time `form:"createAtEnd"  search:"type:lte;column:create_at;table:data_dato" comment:"t in mongoDB"`
	DataDatoOrder
}

type DataDatoOrder struct {
	//inOrder进行绑定，不存在，存在则要求类型为string
	IdOrder        string `form:"idOrder" search:"type:order;column:id;table:data_dato"`
	CreatedAtOrder string `form:"createdAtOrder"  search:"type:order;column:created_at;table:data_dato"`
}

func (m *DataDatoGetPageReq) GetNeedSearch() interface{} {
	return *m
}

// DataDatoGetReq 功能获取请求参数
type DataDatoGetReq struct {
	Id int `uri:"id"`
}

func (s *DataDatoGetReq) GetId() interface{} {
	return s.Id
}

// DataDatoDeleteReq 功能删除请求参数
type DataDatoDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *DataDatoDeleteReq) GetId() interface{} {
	return s.Ids
}
