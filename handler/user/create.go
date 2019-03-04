package user

import (
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"apiserver/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// Create 创建一个新用户
func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest

	var err error
	if err = c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		// c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		//return
	}
	u := model.User{
		Username: r.Username,
		Password: r.Password,
	}
	// adminName := c.Param("username")
	// log.Infof("URL username: %s", adminName)

	// desc := c.Query("desc")
	// log.Infof("URL key param desc: %s", desc)
	// contentType := c.GetHeader("Content-Type")
	// log.Infof("Header Content-Type: %s", contentType)

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)
	// if r.Username == "" {
	// 	err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx"))
	// 	SendResponse(c, err, nil)
	// 	log.Errorf(err, "Get an error")
	// }

	// if errno.IsErrUserNotFound(err) {
	// 	log.Debug("err type is ErrUserNotFound")
	// }

	// if r.Password == "" {
	// 	err = fmt.Errorf("password is empty")
	// 	SendResponse(c, err, nil)
	// }

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	// Insert the user to the database.
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, err)
		return
	}
	rsp := CreateResponse{
		Username: r.Username,
	}
	// Show the user information.
	SendResponse(c, nil, rsp)
	// code, message := errno.DecodeErr(err)
	// c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
