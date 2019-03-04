package user

import (
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Index Get gets an user by the user identifier.
func Index(c *gin.Context) {
	// Get the user by the `username` from the database.
	userID, _ := strconv.Atoi(c.Param("id"))
	user, err := model.Index(uint64(userID))
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
	}

	SendResponse(c, nil, user)
}

// GetUser Get gets an user by the user name.
func GetUser(c *gin.Context) {
	// Get the user by the `username` from the database.
	user, err := model.GetUser(c.Param("username"))
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
	}

	SendResponse(c, nil, user)
}
