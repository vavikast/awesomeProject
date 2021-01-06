package user

import (
	"awesomeProject/apiserver/demo4/handler"
	"awesomeProject/apiserver/demo4/model"
	"awesomeProject/apiserver/demo4/pkg/auth"
	"awesomeProject/apiserver/demo4/pkg/errno"
	"awesomeProject/apiserver/demo4/pkg/token"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// Binding the data with the user struct.
	var u model.UserModel
	if err := c.Bind(&u); err != nil {

		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// Get the user information by the login username.
	fmt.Println(u.Username)
	d, err := model.GetUser(u.Username)
	fmt.Println("JustTest", d)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	// Compare the login password with the user password.
	if err := auth.Compare(d.Password, u.Password); err != nil {
		handler.SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// Sign the json web token.
	t, err := token.Sign(c, token.Context{ID: d.Id, Username: d.Username}, "")
	if err != nil {
		handler.SendResponse(c, errno.ErrToken, nil)
		return
	}

	handler.SendResponse(c, nil, model.Token{Token: t})
}
