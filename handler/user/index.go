package user

import (
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get gets an user by the user identifier.
func Index(c *gin.Context) {
	// Get the user by the `username` from the database.
	userId, _ := strconv.Atoi(c.Param("id"))
	user, err := model.Index(uint64(userId))
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}
