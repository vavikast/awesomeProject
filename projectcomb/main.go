package main

import (
	"awesomeProject/projectcomb/dao"
	"awesomeProject/projectcomb/router"
	"fmt"
)

func main() {
	if err := dao.Init(); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer dao.Close() // 程序退出关闭数据库连接
	r := router.SetupRouter()
	err := r.Run(":8081")
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
