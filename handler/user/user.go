package user

import (
	"apiserver/model"
)

// CreateRequest 请求结构
type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CreateResponse 响应结构
type CreateResponse struct {
	Username string `json:"username"`
}

type ListRequest struct {
	Offset int `form:"offset" json:"offset"`
	Limit  int `form:"limit" json:"limit"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}
