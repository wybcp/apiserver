package user

import (
	"strconv"

	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Delete delete an user by the user identifier.
func Delete(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	var u model.User
	if err := u.Delete(uint64(userID)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
	}

	SendResponse(c, nil, nil)
}
