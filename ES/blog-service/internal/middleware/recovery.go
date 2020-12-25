package middleware

import (
	"awesomeProject/ES/blog-service/global"
	"awesomeProject/ES/blog-service/pkg/app"
	"awesomeProject/ES/blog-service/pkg/email"
	"awesomeProject/ES/blog-service/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Recovery() gin.HandlerFunc {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf(c, "panic recover err :%v", err)
				defailtMailer.SendEmail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出： 发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息：%v", err),
				)
				if err != nil {
					global.Logger.Panicf(c, "mail.Sender err: %v", err)
				}
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
