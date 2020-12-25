package controlller

import (
	"awesomeProject/projectcomb/dao"
	"awesomeProject/projectcomb/logic"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type ProjectMap struct {
	Name   string
	Values interface{}
}

//根据业务分类名称查询业务分类ID
func GetQuerycfidHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "GetQuerycfid.html", nil)

}
func PostQuerycfidHandler(c *gin.Context) {
	emname := c.PostForm("cfname")
	fmt.Println(emname)
	data, err := dao.Querycfid(emname)
	if err != nil {
		fmt.Println("获取业务分类ID错误：", err)
		c.JSON(405, gin.H{"message": "输入错误"})
		return
	}
	m1 := ProjectMap{
		Name:   "业务分类号：",
		Values: data,
	}
	c.HTML(200, "Postindexsingle.html", m1)
}

//根据业务分类ID查询业务分类名称
func GetQuerycfnameHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "GetQuerycfname.html", nil)
}
func PostQuerycfnameHandler(c *gin.Context) {
	cfid := c.PostForm("cfid")
	fmt.Println(cfid)
	atoi, _ := strconv.Atoi(cfid)
	data, err := dao.Querycfname(atoi)
	if err != nil {
		fmt.Println("获取业务分类名称错误：", err)
		c.JSON(405, gin.H{"message": "输入错误"})
		return
	}
	m1 := ProjectMap{
		Name:   "业务分类名称：",
		Values: data,
	}
	c.HTML(200, "Postindexsingle.html", m1)

}

//根据部门名称查询部门ID
func GetQuerydpidHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "GetQuerydpid.html", nil)
}
func PostQuerydpidHandler(c *gin.Context) {
	dpname := c.PostForm("dpname")
	fmt.Println(dpname)
	data, err := dao.Querydpid(dpname)
	if err != nil {
		fmt.Println("获取部门ID错误：", err)
		c.JSON(405, gin.H{"message": "输入错误"})
		return
	}

	m1 := ProjectMap{
		Name:   "部门ID：",
		Values: data,
	}
	c.HTML(200, "Postindexsingle.html", m1)
}

//根据部门ID查询部门名称
func GetQuerydpnameHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "GetQuerydpname.html", nil)
}
func PostQuerydpnameHandler(c *gin.Context) {
	dpid := c.PostForm("dpid")
	fmt.Println(dpid)
	atoi, _ := strconv.Atoi(dpid)
	data, err := dao.Querydpname(atoi)
	if err != nil {
		fmt.Println("获取部门名称错误：", err)
		c.JSON(405, gin.H{"message": "输入错误"})
		return
	}
	m1 := ProjectMap{
		Name:   "部门名称：",
		Values: data,
	}
	c.HTML(200, "Postindexsingle.html", m1)
}

//根据员工名称查询员工ID
func GetQueryemidHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "GetQueryemid.html", nil)
}
func PostQueryemidHandler(c *gin.Context) {
	emname := c.PostForm("emname")
	fmt.Println(emname)
	data, err := dao.Queryemid(emname)
	if err != nil {
		fmt.Println("获取员工ID错误：", err)
		c.JSON(405, gin.H{"message": "输入错误"})
		return
	}
	m1 := ProjectMap{
		Name:   "员工ID如下",
		Values: data,
	}
	c.HTML(200, "Postindexsingle.html", m1)
}

//根据员工ID查询员工名称
func GetQueryemnameHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "GetQueryemname.html", nil)
}
func PostQueryemnameHandler(c *gin.Context) {
	emid := c.PostForm("emid")
	fmt.Println(emid)
	atoi, _ := strconv.Atoi(emid)
	data, err := dao.Queryemname(atoi)
	if err != nil {
		fmt.Println("获取员工名称错误：", err)
		c.JSON(405, gin.H{"message": "输入错误"})
		return
	}
	m1 := ProjectMap{
		Name:   "员工名称名称：",
		Values: data,
	}
	c.HTML(200, "Postindexsingle.html", m1)
}

//根据流程名称查询流程ID
func GetQuerywfidHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "GetQuerywfid.html", nil)
}
func PostQuerywfidHandler(c *gin.Context) {
	wfname := c.PostForm("wfname")
	fmt.Println(wfname)
	data, err := dao.Querywfid(wfname)
	if err != nil {
		fmt.Println("获取流程ID错误：", err)
		c.JSON(405, gin.H{"message": "输入错误"})
		return
	}
	m1 := ProjectMap{
		Name:   "工作流程ID如下",
		Values: data,
	}
	c.HTML(200, "Postindexsingle.html", m1)
}

//根据流程ID查询流程名称
func GetQuerywfnameHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "GetQuerywfname.html", nil)
}
func PostQuerywfnameHandler(c *gin.Context) {
	wfid := c.PostForm("wfid")
	fmt.Println(wfid)
	atoi, _ := strconv.Atoi(wfid)
	data, err := dao.Querywfname(atoi)
	if err != nil {
		fmt.Println("获取流程步骤名称错误：", err)
		c.JSON(405, gin.H{"message": "输入错误"})
		return
	}
	m1 := ProjectMap{
		Name:   "流程步骤名称：",
		Values: data,
	}
	c.HTML(200, "Postindexsingle.html", m1)
}

//根据流程步骤ID查询流程步骤名称
func GetQuerywsidHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "GetQuerywsid.html", nil)
}
func PostQuerywsidHandler(c *gin.Context) {
	var wfidornameinterface interface{}
	wfidorname := c.PostForm("wfidorname")
	wsname := c.PostForm("wsname")
	wfidornameinterface = wfidorname
	if logic.IsNum(wfidorname) {
		wfidornameint, err := strconv.Atoi(wfidorname)
		if err != nil {
			fmt.Println("数字转换错误")
			return
		}
		wfidornameinterface = wfidornameint
	}
	data, err := dao.Querywsid(wfidornameinterface, wsname)
	if err != nil {
		fmt.Println("获取流程步骤ID错误：", err)
		c.JSON(405, gin.H{"message": "输入错误"})
		return
	}
	m1 := ProjectMap{
		Name:   "流程步骤ID如下",
		Values: data,
	}
	c.HTML(200, "Postindexsingle.html", m1)
}

//根据流程步骤名称查询流程步骤ID
func GetQuerywsnameHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "GetQuerywsname.html", nil)
}
func PostQuerywsnameHandler(c *gin.Context) {
	var wfidornameinterface interface{}
	var wsidint int
	wfidorname := c.PostForm("wfidorname")
	wsid := c.PostForm("wsid")
	wfidornameinterface = wfidorname
	if logic.IsNum(wfidorname) {
		wfidornameint, err := strconv.Atoi(wfidorname)
		if err != nil {
			fmt.Println("数字转换错误")
			return
		}
		wfidornameinterface = wfidornameint
	}
	if logic.IsNum(wsid) {
		d, err := strconv.Atoi(wsid)
		wsidint = d
		if err != nil {
			fmt.Println("数字转换错误")
			return
		}
	}
	data, err := dao.Querywsname(wfidornameinterface, wsidint)
	if err != nil {
		fmt.Println("获取流程步骤名称错误：", err)
		c.JSON(405, gin.H{"message": "输入错误"})
		return
	}

	m1 := ProjectMap{
		Name:   "流程步骤名称如下：",
		Values: data,
	}
	c.HTML(200, "Postindexsingle.html", m1)
}

//根据流程名查询完整流程步骤
func GetQuerywsHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "GetQueryws.html", nil)
}
func PostQuerywsHandler(c *gin.Context) {
	var wfidornameinterface interface{}
	wfidorname := c.PostForm("wfidorname")
	wfidornameinterface = wfidorname
	if logic.IsNum(wfidorname) {
		wfidornameint, err := strconv.Atoi(wfidorname)
		if err != nil {
			fmt.Println("数字转换错误")
			return
		}
		wfidornameinterface = wfidornameint
	}
	fmt.Println(wfidorname, "one", "two")
	data, err := dao.Queryws(wfidornameinterface)
	if err != nil {
		fmt.Println("查询流程错误", err)
		c.JSON(405, gin.H{"message": "输入错误"})
		return
	}
	m1 := ProjectMap{
		Name:   "完整工作流程如下:",
		Values: data,
	}
	c.HTML(200, "Postindex.html", m1)
}

//根据流程名查询授权人，经办人，主办人
func GetQuerypersonHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "GetQueryperson.html", nil)
}
func PostQuerypersonHandler(c *gin.Context) {
	var wfidornameint, wsidornameint, flagint int
	var err error
	wfidorname := c.PostForm("wfidorname")
	wsidorname := c.PostForm("wsidorname")
	flag := c.PostForm("flag")
	fmt.Println(wfidorname, wsidorname, flag)

	if logic.IsNum(wfidorname) {
		wfidornameint, err = strconv.Atoi(wfidorname)
		if err != nil {
			fmt.Println("数字转换错误")
			return
		}
	} else {
		wfidornameint, err = dao.Querywfid(wfidorname)
		if err != nil {
			fmt.Println("查询流程ID操作", err)
			return
		}
	}
	if logic.IsNum(wsidorname) {
		wsidornameint, err = strconv.Atoi(wsidorname)
		if err != nil {
			fmt.Println("数字转换错误")
			return
		}
	} else {
		wsidornameint, err = dao.Querywsid(wfidorname, wsidorname)
		if err != nil {
			fmt.Println("查询流程ID操作", err)
			return
		}
	}
	if logic.IsNum(flag) {
		flagint, err = strconv.Atoi(flag)
		if err != nil {
			fmt.Println("数字转换错误")
			return
		}
	}

	data, err := dao.QueryPerson(wfidornameint, wsidornameint, flagint)
	if err != nil {
		fmt.Println("查询员工所在流程错误：", err)
		c.JSON(405, gin.H{"message": "输入错误"})
		return
	}
	if flagint == 0 {

		m1 := ProjectMap{
			Name:   "授权人如下:",
			Values: data,
		}
		c.HTML(200, "Postindex.html", m1)
	}
	if flagint == 1 {
		m1 := ProjectMap{
			Name:   "经办人如下:",
			Values: data,
		}
		c.HTML(200, "Postindex.html", m1)
	}
	if flagint == 2 {
		m1 := ProjectMap{
			Name:   "主办人如下:",
			Values: data,
		}
		c.HTML(200, "Postindex.html", m1)
	}
}

//根据需要查询全部部门，全部员工，全部工作流
func GetQueryallHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "GetQueryall.html", nil)
}
func PostQueryallHandler(c *gin.Context) {
	var err error
	flag := c.PostForm("flag")
	fmt.Println(flag)
	data, err := dao.QueryAll(flag)
	if err != nil {
		fmt.Println("获取员工信息错误：", err)
		c.JSON(405, gin.H{"message": "输入错误"})
		return
	}
	if strings.Contains(flag, "dp") {
		m1 := ProjectMap{
			Name:   "全部部门如下:",
			Values: data,
		}
		c.HTML(200, "Postindex.html", m1)
	}
	if strings.Contains(flag, "em") {
		m1 := ProjectMap{
			Name:   "全部员工如下:",
			Values: data,
		}
		c.HTML(200, "Postindex.html", m1)
	}
	if strings.Contains(flag, "wf") {
		m1 := ProjectMap{
			Name:   "全部工作流如下:",
			Values: data,
		}
		c.HTML(200, "Postindex.html", m1)
	}
	if strings.Contains(flag, "cf") {
		m1 := ProjectMap{
			Name:   "全部流程分类如下:",
			Values: data,
		}
		c.HTML(200, "Postindex.html", m1)
	}

}

//查询含有判断条件的流程
func GetQueryjudgeHandler(c *gin.Context) {
	data, err := dao.Queryjudge()
	if err != nil {
		fmt.Println("获取判断流程操作：", err)
		c.JSON(405, gin.H{"message": "输入错误"})
		return
	}
	m1 := ProjectMap{
		Name:   "含有判断的流程如下:",
		Values: data,
	}
	c.HTML(200, "Postindex.html", m1)
	//c.HTML(http.StatusOK, "GetQueryjudge.html", nil)
}
func PostQueryjudgeHandler(c *gin.Context) {
	data, err := dao.Queryjudge()
	if err != nil {
		fmt.Println("获取判断流程操作：", err)
		c.JSON(405, gin.H{"message": "输入错误"})
		return
	}
	m1 := ProjectMap{
		Name:   "含有判断的流程如下:",
		Values: data,
	}
	c.HTML(200, "Postindex.html", m1)
}

//查询员工所在流程以及流程步骤
func GetQuerymatchwfHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "GetQuerymatchwf.html", nil)
}
func PostQuerymatchwfHandler(c *gin.Context) {
	var flagint int
	emname := c.PostForm("emname")
	flag := c.PostForm("flag")
	fmt.Println(emname, flag)
	if logic.IsNum(flag) {
		d, err := strconv.Atoi(flag)
		flagint = d
		if err != nil {
			fmt.Println("数字转换错误")
			return
		}
	}
	data, err := dao.QueryMatchWf(emname, flagint)
	if err != nil {
		fmt.Println("查询流程信息错误：", err)
		c.JSON(405, gin.H{"message": "输入错误"})
		return
	}
	if flagint == 0 {
		m1 := ProjectMap{
			Name:   "经办人所在流程如下:",
			Values: data,
		}
		c.HTML(200, "Postindex.html", m1)
	}
	if flagint == 1 {
		m1 := ProjectMap{
			Name:   "主办人所在流程如下:",
			Values: data,
		}
		c.HTML(200, "Postindex.html", m1)
	}

}
