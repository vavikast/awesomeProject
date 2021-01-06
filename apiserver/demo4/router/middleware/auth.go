package middleware

import (
	"awesomeProject/apiserver/demo4/handler"
	"awesomeProject/apiserver/demo4/pkg/errno"
	"awesomeProject/apiserver/demo4/pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
