package main

import (
	"awesomeProject/ES/chatroom/server"
	"fmt"
	"github.com/labstack/gommon/log"
	"net/http"
)


var (
	addr   = ":2022"
	banner = `
    ____              _____
   |    |    |   /\     |
   |    |____|  /  \    | 
   |    |    | /----\   |
   |____|    |/      \  |
Go语言编程之旅 —— 一起用Go做项目：ChatRoom，start on：%s
`
)

func main()  {
	fmt.Printf(banner+"\n",addr)
	server.RegisterHandler()
	log.Fatal(http.ListenAndServe(addr,nil))
}