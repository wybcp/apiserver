package user

import (
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service"
	"github.com/gin-gonic/gin"
)

// List list the users in the database.
func List(c *gin.Context) {
	var r ListRequest
	//使用默认的bind ，get方法只处理 form 格式请求
	//if err := c.Bind(&r); err != nil {
	//使用BindJSON ，get方法可以处理 json 格式
	if err := c.BindJSON(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	users, count, err := service.ListUser(r.Offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   users,
	})
}
