package user

import (
	"awesomeProject/apiserver/demo4/handler"
	"awesomeProject/apiserver/demo4/model"
	"awesomeProject/apiserver/demo4/pkg/errno"
	"awesomeProject/apiserver/demo4/util"
	"github.com/gin-gonic/gin"
	"github.com/zxmrlc/log"
	"github.com/zxmrlc/log/lager"
)

func Create(c *gin.Context) {
	log.Info("User create funttion called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}
	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	rsp := CreateResponse{
		Username: r.Username,
	}
	// Show the user information.
	handler.SendResponse(c, nil, rsp)

}
