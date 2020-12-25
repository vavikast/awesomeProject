package router

import (
	"awesomeProject/projectcomb/controlller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {

	r := gin.New()
	//r.LoadHTMLFiles("./projectcomb/templates/index.html")
	r.LoadHTMLGlob("projectcomb/templates/*.html")
	////r.LoadHTMLGlob("templates/**/*")
	////r.StaticFS("/templates", http.Dir("./projectcomb/templates/"))
	//r.LoadHTMLGlob("./projectcomb/templates/**/*")
	//r.StaticFS("./projectcomb/templates", http.Dir("./"))
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//根据业务分类名称查询业务分类ID
	r.GET("/querycfid", controlller.GetQuerycfidHandler)
	r.POST("/querycfid", controlller.PostQuerycfidHandler)
	//根据业务分类ID查询业务分类名称
	r.GET("/querycfname", controlller.GetQuerycfnameHandler)
	r.POST("/querycfname", controlller.PostQuerycfnameHandler)
	//根据部门名称查询部门ID
	r.GET("/querydpid", controlller.GetQuerydpidHandler)
	r.POST("/querydpid", controlller.PostQuerydpidHandler)
	//根据部门ID查询部门名称
	r.GET("/querydpname", controlller.GetQuerydpnameHandler)
	r.POST("/querydpname", controlller.PostQuerydpnameHandler)
	//根据员工名称查询员工ID
	r.GET("/queryemid", controlller.GetQueryemidHandler)
	r.POST("/queryemid", controlller.PostQueryemidHandler)
	//根据员工ID查询员工名称
	r.GET("/queryemname", controlller.GetQueryemnameHandler)
	r.POST("/queryemname", controlller.PostQueryemnameHandler)
	//根据流程名称查询流程ID
	r.GET("/querywfid", controlller.GetQuerywfidHandler)
	r.POST("/querywfid", controlller.PostQuerywfidHandler)
	//根据流程ID查询流程名称
	r.GET("/querywfname", controlller.GetQuerywfnameHandler)
	r.POST("/querywfname", controlller.PostQuerywfnameHandler)
	//根据流程步骤ID查询流程步骤名称
	r.GET("/querywsid", controlller.GetQuerywsidHandler)
	r.POST("/querywsid", controlller.PostQuerywsidHandler)
	//根据流程步骤名称查询流程步骤ID
	r.GET("/querywsname", controlller.GetQuerywsnameHandler)
	r.POST("/querywsname", controlller.PostQuerywsnameHandler)
	//根据流程名查询完整流程步骤
	r.GET("/queryws", controlller.GetQuerywsHandler)
	r.POST("/queryws", controlller.PostQuerywsHandler)
	//根据流程名查询授权人，经办人，主办人
	r.GET("/queryperson", controlller.GetQuerypersonHandler)
	r.POST("/queryperson", controlller.PostQuerypersonHandler)
	//根据需要查询全部部门，全部员工，全部工作流
	r.GET("/queryall", controlller.GetQueryallHandler)
	r.POST("/queryall", controlller.PostQueryallHandler)
	//查询含有判断条件的流程
	r.GET("/queryjudge", controlller.GetQueryjudgeHandler)
	r.POST("/queryjudge", controlller.PostQueryjudgeHandler)
	//查询员工所在流程以及流程步骤
	r.GET("/querymatchwf", controlller.GetQuerymatchwfHandler)
	r.POST("/querymatchwf", controlller.PostQuerymatchwfHandler)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
