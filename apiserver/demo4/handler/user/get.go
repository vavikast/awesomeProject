package user

import (
	"awesomeProject/apiserver/demo4/handler"
	"awesomeProject/apiserver/demo4/model"
	"awesomeProject/apiserver/demo4/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	username := c.Param("username")
	user, err := model.GetUser(username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	handler.SendResponse(c, nil, user)
}
