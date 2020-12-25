package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

//聊天室
//主要功能： 上线，下线，全员通知
//单对单聊天
type  User struct {
	ID int
	Addr string
	EnterAt string
	MessageChannel chan string

}


func (u *User)String()  string {
	return "this is "+strconv.Itoa(u.ID)+": "+u.Addr

}
var  IdList map[int]int

var (
	enteringChannel = make(chan *User)
	leavingChannel = make(chan *User)
	GlobalMessage = make(chan string,1024)
	Userlist = make(map[*User]struct{})
)

func GenerateID() int {
	if  IdList == nil{
		IdList = make(map[int]int)
	}
	rand.Seed(time.Now().Unix())
	n := rand.Intn(20000000)
	if  IdList[n] == 1{
		return  GenerateID()
	}
	IdList[n] = 1
	return  n
}


func main()  {
	// 成员信息：(名称，通信内容，时间，mux)
	// 开启服务
	listener, err := net.Listen("tcp", ":8020")
	if err != nil {
		panic(err)
	}

	go broadcaster()
	// 处理数据
	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println(err)
			continue
		}
		// 数据通信
		fmt.Println("关联数据")
		go handleConn(conn)
	}






	//成员管理

}

func broadcaster() {
	for {
		select {
		case user := <-enteringChannel:
			Userlist[user] = struct {}{}
		case msg := <-GlobalMessage:
			for i,_:= range Userlist{
				i.MessageChannel <-msg
			}
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	user := &User{
		ID:             GenerateID(),
		Addr:           conn.RemoteAddr().String(),
		EnterAt:        time.Now().Format("2006-01-02 15:04:06"),
		MessageChannel: make(chan  string,1024),
	}
	fmt.Println(user.ID)
	user.MessageChannel <-  "Yes It's me"+strconv.Itoa(user.ID)

	message := <- user.MessageChannel
	GlobalMessage <- message
	fmt.Println(message)
	s := user.String()+message
	enteringChannel <- user
	fmt.Println(s)
	go sendMessage(conn,s)

	input := bufio.NewScanner(conn)
	for input.Scan(){
		words := strconv.Itoa(user.ID)+":"+input.Text()
		fmt.Println(words)
	}
	if err := input.Err();err !=nil{
		log.Println("读取错误：",err)
	}




}
func sendMessage(conn net.Conn,s string)  {
	fmt.Fprintln(conn,s)

}