package main

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	_ = iota
	CF
	DP
	EM
	WF
	WS
)

type Classifi struct {
	Nuid   int       `json:"nuid" db:"nuid"`
	Cfid   int       `json:"cfid" db:"cfid"`
	Cfname string    `json:"cfname" db:"cfname"`
	Cftime time.Time `json:"cftime" db:"cftime"`
}
type Department struct {
	Nuid   int    `json:"nuid" db:"nuid"`
	Dpid   int    `json:"dpid" db:"dpid"`
	Dpname string `json:"dpname" db:"dpname"`
}
type Employee struct {
	Nuid   int    `json:"nuid" db:"nuid"`
	Emid   int    `json:"emid" db:"emid"`
	Emname string `json:"emname" db:"emname"`
	Emmale int    `json:"emmale" db:"emmale"`
	Dpid   int    `json:"dpid" db:"dpid"`
}
type Workflow struct {
	Nuid   int       `json:"nuid" db:"nuid"`
	Wfid   int       `json:"wfid" db:"wfid"`
	Wfname string    `json:"wfname" db:"wfname"`
	Cfid   int       `json:"cfid" db:"cfid"`
	Wftime time.Time `json:"wftime"  db:"wftime"`
}
type Workstep struct {
	Nuid         int    `json:"nuid" db:"nuid"`
	Wsid         int    `json:"wsid" db:"wsid"`
	Wsname       string `json:"wsname" db:"wsname"`
	Wfid         int    `json:"wfid" db:"wfid"`
	Wsjudge      int    `json:"wsjudge" db:"wsjudge"`
	Wsfork       int    `json:"wsfork" db:"wsfork"`
	Wsmerge      int    `json:"wsmerge" db:"wsmerge"`
	Wspreid      string `json:"wspreid" db:"wspreid"`
	Wsposid      string `json:"wsposid" db:"wsposid"`
	Wsauthorizer string `json:"wsauthorizer" db:"wsauthorizer"`
	Wsoperator   string `json:"wsoperator" db:"wsoperator"`
	Wsorganiser  string `json:"wsorganiser" db:"wsorganiser"`
}

var db *sqlx.DB

type User struct {
	Id   int
	Name string
	Age  int
}

func initDB() (err error) {
	dsn := "wl:acfi1e@ije#201@tcp(192.168.20.201:3306)/workflows?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

// 查询多条数据示例
// 查询单条数据示例

func main() {
	err := initDB()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(CF)
	fmt.Println(DP)
	fmt.Println(EM)
	fmt.Println(WF)
	fmt.Println(WS)
	// 查询流程分类名称
	querycfname, err := Querycfname(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(querycfname)
	// 查询流程分类id
	querycfid, err := Querycfid("人事类")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(querycfid)

	// 查询部门名称
	querydpname, err := Querydpname(1002)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(querydpname)
	// 查询部门id
	querydpid, err := Querydpid("综合管理中心")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(querydpid)

	// 查询员工名称
	queryemname, err := Queryemname(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(queryemname)
	// 查询员工id
	queryemid, err := Queryemid("王阳")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(queryemid)

	// 查询工作流名称
	querywfname, err := Querywfname(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(querywfname)
	// 查询工作流程id
	querywfid, err := Querywfid("用印申请表")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(querywfid)

	// 查询流程步骤名称
	querywsname, err := Querywsname(1, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(querywsname)
	// 查询流程步骤id
	querywsid, err := Querywsid(20, "中心领导审批")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(querywsid)

	// 根据流程名查询整流程
	ws, err := Queryws("用印申请表")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ws)
	// 根据流程编号查询整流程
	ws, err = Queryws(18)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ws)

	//根据流程名查询授权人0，经办人1，主办人2
	person, err := QueryPerson(18, 3, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(person)

	//根据员工和部门id查询名称
	emdpname, err := Queryemdpname(17)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("通过员工和部门id查询名称:" + emdpname)
	emdpname1, err := Queryemdpname(1000)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("通过员工和部门id查询名称1:" + emdpname1)

	//根据员工和部门名称查询id
	emdpid, err := Queryemdpid("王阳")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("通过员工和部门名称查询id1:", emdpid)
	emdpid2, err := Queryemdpid("综合管理中心")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("通过员工和部门名称查询id2:", emdpid2)

	//查询全部业务分类，查询全部部门，查询全部员工，查询全部工作流 cf,dp,em,wf,ws
	allcf, err := QueryAll("cf")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("查询结果是：", allcf)
	alldp, err := QueryAll("dp")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("查询结果是：", alldp)
	allwf, err := QueryAll("wf")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("查询结果是：", allwf)
	allem, err := QueryAll("em")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("查询结果是：", allem)

	//查询所有含有判断的流程
	queryjudge, err := Queryjudge()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("查询结果是：", queryjudge)
	//查询判断从办人所在节点
	wf, err := QueryMatchWfoperator("黄干飞")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("查询结果是：", wf)
	wf1, err := QueryMatchWforganiser("黄干飞")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("查询结果是：", wf1)
}

//查询流程分类名称
func Querycfname(cfid int) (cfname string, err error) {
	sqlStr := "select cfname from classifi where cfid = ? "
	var cf Classifi
	err = db.Get(&cf, sqlStr, cfid)
	if err != nil {
		fmt.Printf("流程分类名称获取失败, err: %v\n", err)
		return "", err
	}
	return cf.Cfname, nil
}

//查询流程分类编号
func Querycfid(cfname string) (cfid int, err error) {
	sqlStr := "select cfid from classifi where cfname =  ? "
	var cf Classifi
	err = db.Get(&cf, sqlStr, cfname)
	if err != nil {
		fmt.Printf("流程分类编号获取失败, err: %v\n", err)
		return cfid, err
	}
	return cf.Cfid, nil
}

//查询部门名称
func Querydpname(dpid int) (dpname string, err error) {
	sqlStr := "select dpname from department where dpid = ? "
	var dp Department
	err = db.Get(&dp, sqlStr, dpid)
	if err != nil {
		fmt.Printf("部门名称获取失败, err: %v\n", err)
		return "", err
	}
	return dp.Dpname, nil
}

//查询部门id
func Querydpid(dpname string) (dpid int, err error) {
	sqlStr := "select dpid from department where dpname = ? "
	var dp Department
	err = db.Get(&dp, sqlStr, dpname)
	if err != nil {
		fmt.Printf("部门id获取失败, err: %v\n", err)
		return dpid, err
	}
	return dp.Dpid, nil
}

//查询员工名称
func Queryemname(emid int) (emname string, err error) {
	sqlStr := "select emname from employee where emid = ? "
	var em Employee
	err = db.Get(&em, sqlStr, emid)
	if err != nil {
		fmt.Printf("员工名称获取失败, err: %v\n", err)
		return "", err
	}
	return fmt.Sprint(em.Emname), nil
}

//查询员工id
func Queryemid(emname string) (emid int, err error) {
	sqlStr := "select emid from employee where emname = ? "
	var em Employee
	err = db.Get(&em, sqlStr, emname)
	if err != nil {
		fmt.Printf("员工id获取失败, err: %v\n", err)
		return emid, err
	}
	return em.Emid, nil
}

//查询工作流名称
func Querywfname(wfid int) (wfname string, err error) {
	sqlStr := "select wfname from workflow where wfid = ? "
	var wf Workflow
	err = db.Get(&wf, sqlStr, wfid)
	if err != nil {
		fmt.Printf("工作流程名称获取失败, err: %v\n", err)
		return "", err
	}
	return fmt.Sprint(wf.Wfname), nil
}

//查询工作流id
func Querywfid(wfname string) (wfid int, err error) {
	sqlStr := "select wfid from workflow where wfname= ? "
	var wf Workflow
	err = db.Get(&wf, sqlStr, wfname)
	if err != nil {
		fmt.Printf("工作流程id获取失败, err: %v\n", err)
		return wfid, err
	}
	return wf.Wfid, nil
}

//查询流程步骤名称
func Querywsname(wfid, wsid int) (wsname string, err error) {
	sqlStr := "select wsname from workstep where wfid=? and wsid = ? "
	var ws Workstep
	err = db.Get(&ws, sqlStr, wfid, wsid)
	if err != nil {
		fmt.Printf("流程步骤名称获取失败, err: %v\n", err)
		return "", err
	}
	return fmt.Sprint(ws.Wsname), nil
}

//查询流程步骤id
func Querywsid(wfid int, wsname string) (wsid int, err error) {
	sqlStr := "select wsid from workstep where wfid=? and wsname = ? "
	var ws Workstep
	err = db.Get(&ws, sqlStr, wfid, wsname)
	if err != nil {
		fmt.Printf("流程步骤名称获取失败, err: %v\n", err)
		return wsid, err
	}
	return ws.Wsid, nil
}

//查询特定工作流程步骤
func Queryws(qr interface{}) (value interface{}, err error) {
	steps := make([]string, 0)
	v := reflect.ValueOf(qr)
	k := v.Kind()
	switch k {
	case reflect.Int:
		wss, err := queryMultiws(qr.(int))
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return nil, err
		}
		for _, ws := range wss {
			wsone := fmt.Sprint(ws.Wsid, ".", ws.Wsname)
			steps = append(steps, wsone)
		}
		return steps, nil
	case reflect.String:
		sqlStr := "select wfid from workflow where  wfname = ? "
		var ws Workstep
		err = db.Get(&ws, sqlStr, qr.(string))
		if err != nil {
			fmt.Printf("整流程步骤名称获取失败, err: %v\n", err)
			return nil, err
		}
		wss, err := queryMultiws(ws.Wfid)
		if err != nil {
			fmt.Printf("mutiquery failed, err:%v\n", err)
			return nil, err
		}
		for _, ws := range wss {
			wsone := fmt.Sprint(ws.Wsid, ".", ws.Wsname)
			steps = append(steps, wsone)
		}
		return steps, nil
	default:
		return nil, errors.New("输入类型错误")
	}
	return nil, errors.New("输入类型错误")
}

// 查询多条数据示例
func queryMultiws(wsid int) (wsname []Workstep, err error) {
	sqlStr := "select wsid,wsname from workstep where wfid=  ?"
	var wss []Workstep
	err = db.Select(&wss, sqlStr, wsid)
	if err != nil {
		fmt.Printf("multiquery failed, err:%v\n", err)
		return nil, err
	}
	return wss, nil
}

//根据流程id查询wsauthorizer(授权人),wsoperator(经办人),wsorganiser(主办人)
func QueryPerson(wfid, wsid, flag int) (all []string, err error) {
	all = make([]string, 0)
	sqlStr := "select wsauthorizer,wsoperator,wsorganiser from workstep where wfid=  ? and wsid= ?"
	var ws Workstep
	err = db.Get(&ws, sqlStr, wfid, wsid)
	if err != nil {
		fmt.Printf("部门查询错误, err:%v\n", err)
		return nil, err
	}
	if flag == 0 {
		wsauthorizer := ws.Wsauthorizer
		wsau := strings.Split(wsauthorizer, ",")
		for _, au := range wsau {
			nu, _ := strconv.Atoi(au)
			emname, err := Queryemname(nu)
			if err != nil {
				fmt.Printf("员工id查询错误, err:%v\n", err)
				return nil, err
			}
			all = append(all, emname)
		}
		return all, nil
	}
	if flag == 1 {
		wsoperator := ws.Wsoperator
		wsop := strings.Split(wsoperator, ",")
		for _, op := range wsop {
			nu, _ := strconv.Atoi(op)
			emname, err := Queryemname(nu)
			if err != nil {
				fmt.Printf("员工id查询错误, err:%v\n", err)
				return nil, err
			}
			all = append(all, emname)
		}
		return all, nil

	}
	if flag == 2 {
		wsorganiser := ws.Wsorganiser
		wsor := strings.Split(wsorganiser, ",")
		for _, or := range wsor {
			nu, _ := strconv.Atoi(or)
			emname, err := Queryemname(nu)
			if err != nil {
				fmt.Printf("员工id查询错误, err:%v\n", err)
				return nil, err
			}
			all = append(all, emname)
		}
		return all, nil

	}
	return nil, errors.New("部门查询结果错误")
}

//通过id查询员工和部门名称
func Queryemdpname(emdpid int) (emdpname string, err error) {
	if emdpid < 1000 {
		queryemdpname, err := Queryemname(emdpid)
		if err != nil {
			fmt.Printf("员工名称获取失败, err: %v\n", err)
			return emdpname, err
		}
		return queryemdpname, nil
	} else {
		querydpname, err := Querydpname(emdpid)
		if err != nil {
			fmt.Printf("部门名称获取失败, err: %v\n", err)
			return emdpname, err
		}
		return querydpname, err
	}
	return emdpname, nil
}

//通过员工和部门名称查询员工和部门id
//查询员工id
func Queryemdpid(emdpname string) (emdpid int, err error) {
	queryemid, err := Queryemid(emdpname)
	if err != nil {
		querydpid, err := Querydpid(emdpname)
		if err != nil {
			fmt.Printf("部门id获取失败, err: %v\n", err)
			return emdpid, err
		}
		return querydpid, nil
	}
	return queryemid, nil

}

//查询全部业务分类，查询全部部门，查询全部员工，查询全部工作流
func QueryAll(opt string) (list []string, err error) {
	list = make([]string, 0)
	if strings.Contains(opt, "cf") {
		sqlStr := "select cfid, cfname from classifi where cfid > ?"
		var cfs []Classifi
		err := db.Select(&cfs, sqlStr, 0)
		if err != nil {
			fmt.Printf("all-qf-query failed, err:%v\n", err)
			return nil, err
		}
		for _, cf := range cfs {
			cfbind := strconv.Itoa(cf.Cfid) + "." + cf.Cfname
			list = append(list, cfbind)
		}
		return list, nil
	}
	if strings.Contains(opt, "dp") {
		sqlStr := "select dpid, dpname from department where dpid > ?"
		var dps []Department
		err := db.Select(&dps, sqlStr, 0)
		if err != nil {
			fmt.Printf("all-dps-query failed, err:%v\n", err)
			return nil, err
		}
		for _, dp := range dps {
			dpbind := strconv.Itoa(dp.Dpid) + "." + dp.Dpname
			list = append(list, dpbind)
		}
		return list, nil
	}
	if strings.Contains(opt, "em") {
		sqlStr := "select emid, emname from employee where emid > ?"
		var ems []Employee
		err := db.Select(&ems, sqlStr, 0)
		if err != nil {
			fmt.Printf("all-ems-query failed, err:%v\n", err)
			return nil, err
		}
		for _, em := range ems {
			embind := strconv.Itoa(em.Emid) + "." + em.Emname
			list = append(list, embind)
		}
		return list, nil
	}
	if strings.Contains(opt, "wf") {
		sqlStr := "select wfid, wfname from workflow where wfid > ?"
		var wfs []Workflow
		err := db.Select(&wfs, sqlStr, 0)
		if err != nil {
			fmt.Printf("all-wfs-query failed, err:%v\n", err)
			return nil, err
		}
		for _, wf := range wfs {
			wfbind := strconv.Itoa(wf.Wfid) + "." + wf.Wfname
			list = append(list, wfbind)
		}
		return list, nil
	}
	return nil, errors.New("输入错误，请正确输入cf(业务流程分类),dp(部门),em(员工),wf(工作流)")

}

//查询判断所在节点
func Queryjudge() (wfnames []string, err error) {
	sqlStr := "select wfid from workstep where wsjudge=  ?"
	var wss []Workstep
	wfnames = make([]string, 0)
	err = db.Select(&wss, sqlStr, 1)
	if err != nil {
		fmt.Printf("multiquery failed, err:%v\n", err)
		return wfnames, err
	}
	for _, ws := range wss {
		querywfname, err := Querywfname(ws.Wfid)
		if err != nil {
			fmt.Printf("queryjudge querywfname error, err:%v\n", err)
			return wfnames, err
		}
		wfname := fmt.Sprint(ws.Wfid, ".", querywfname)
		wfnames = append(wfnames, wfname)

	}
	return wfnames, nil
}

//通过人查归属流程+排除开始一个项目。（wsoperator+wsorganiser）只有这两个项目是关键，反推（本例的重点）
//基于现状，现在只关注wsoperator主办人和wsorganiser从办人。 从办人里面一定包含主办人。
//以公司和部门为单位的流程，暂时不在考虑范围。1.工作流程中暂时没有部门为单位的流程节点 2.以公司为单位的流程压根都不需要查询。
//1.先把查询人更换成查询员工id、
//2. 将所有数据库中的从办导出来，形成nuid,wsoperator|wsauthorizer结构。
//3.遍历里面的程序，如果包含此人，则显示流程名称和节点。

//主办人查询
func QueryMatchWfoperator(emname string) (matchwfs []string, err error) {
	emid, err := Queryemdpid(emname)
	matchwfs = make([]string, 0)
	if err != nil {
		fmt.Println("queryMatch , queryemdpidName error")
		return nil, err
	}
	sqlStr := "select nuid,wfid,wsid,wsname,wsoperator from workstep where wfid > ?"
	var wfs []Workstep
	err = db.Select(&wfs, sqlStr, 0)
	if err != nil {
		fmt.Printf("all-wfs-query failed, err:%v\n", err)
		return nil, err
	}

	for _, wf := range wfs {
		if strings.Contains(wf.Wsoperator, strconv.Itoa(emid)) {
			splits := strings.Split(wf.Wsoperator, ",")
			for _, split := range splits {
				if strings.EqualFold(split, strconv.Itoa(emid)) {
					querywfname, err := Querywfname(wf.Wfid)
					if err != nil {
						fmt.Printf("querymatchWf  query wfname failed, err:%v\n", err)
						return nil, err
					}
					matchwf := fmt.Sprint(wf.Wfid, ".", querywfname, "-", wf.Wsname)
					matchwfs = append(matchwfs, matchwf)
				}
			}
		}

	}
	return matchwfs, nil

}

//从办人查询
func QueryMatchWforganiser(emname string) (matchwfs []string, err error) {
	emid, err := Queryemdpid(emname)
	matchwfs = make([]string, 0)
	if err != nil {
		fmt.Println("queryMatchwforganiser , queryemdpidName error")
		return nil, err
	}
	sqlStr := "select nuid,wfid,wsid,wsname,wsorganiser from workstep where wfid > ?"
	var wfs []Workstep
	err = db.Select(&wfs, sqlStr, 0)
	if err != nil {
		fmt.Printf("queryMatchwforganiser, err:%v\n", err)
		return nil, err
	}

	for _, wf := range wfs {
		if strings.Contains(wf.Wsorganiser, strconv.Itoa(emid)) {
			splits := strings.Split(wf.Wsorganiser, ",")
			for _, split := range splits {
				if strings.EqualFold(split, strconv.Itoa(emid)) {
					querywfname, err := Querywfname(wf.Wfid)
					if err != nil {
						fmt.Printf("queryMatchwforganiser  query wfname failed, err:%v\n", err)
						return nil, err
					}
					matchwf := fmt.Sprint(wf.Wfid, ".", querywfname, "-", wf.Wsname)
					matchwfs = append(matchwfs, matchwf)
				}
			}
		}

	}
	return matchwfs, nil

}

//前后流程串联+图标
