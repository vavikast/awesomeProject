package api

import (
	"awesomeProject/ES/blog-service/global"
	"awesomeProject/ES/blog-service/internal/service"
	"awesomeProject/ES/blog-service/pkg/app"
	"awesomeProject/ES/blog-service/pkg/convert"
	"awesomeProject/ES/blog-service/pkg/errcode"
	"awesomeProject/ES/blog-service/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}

}
func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c, "svc.UploadFile err : %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})

}
