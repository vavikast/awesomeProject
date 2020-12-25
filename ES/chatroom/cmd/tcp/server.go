package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":2020")
	if err != nil {
		panic(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConn(conn)
	}
}

type User struct {
	ID             int
	Addr           string
	EnterAt        time.Time
	MessageChannel chan string
}

func (u *User) String() string {
	return u.Addr + ",UID:" + strconv.Itoa(u.ID) + ", Enter At:" + u.EnterAt.Format("2006-01-02 15:04:05+8000")
}

//给用户发送信息
type Message struct {
	OwerID  int
	Content string
}

var (
	//新用户到来，通过该channel进行登记
	enteringChannel = make(chan *User)

	//用户离开，通过该channel进行登记
	leavingChannel = make(chan *User)

	//广播专用的用户普通消息channel,缓冲尽可能避免异常情况阻塞
	messageChannel = make(chan Message, 8)
)

//broadcaster记录聊天室用户，并进行广播
func broadcaster() {
	users := make(map[*User]struct{})
	for {
		select {
		case user := <-enteringChannel:
			users[user] = struct{}{}
		case user := <-leavingChannel:
			delete(users, user)
			close(user.MessageChannel)
		case msg := <-messageChannel:
			//给所有在线用户发送消息
			for user := range users {
				if user.ID == msg.OwerID {
					continue
				}
				user.MessageChannel <- msg.Content
			}

		}

	}

}
func handleConn(conn net.Conn) {
	defer conn.Close()

	//新用户进来构建实例
	user := &User{
		ID:             GenUserID(),
		Addr:           conn.RemoteAddr().String(),
		EnterAt:        time.Now(),
		MessageChannel: make(chan string, 8),
	}

	go sendMessage(conn, user.MessageChannel)

	//向当前用户发送欢迎信息，通知用户的到来
	user.MessageChannel <- "Welcome, " + user.String()
	msg := Message{
		OwerID:  user.ID,
		Content: "user:`" + strconv.Itoa(user.ID) + "` has enter",
	}

	messageChannel <- msg

	//将该记录记录到全局用户列表
	enteringChannel <- user

	//按超时踢出
	var userActive = make(chan struct{})
	go func() {
		d := time.Minute * 1
		timer := time.NewTimer(d)
		for {
			select {
			case <-timer.C:
				conn.Close()
			case <-userActive:
				timer.Reset(d)
			}
		}
	}()

	//循环读取用户的输入
	input := bufio.NewScanner(conn)
	for input.Scan() {
		msg.Content = strconv.Itoa(user.ID) + ":" + input.Text()
		messageChannel <- msg
		userActive <- struct{}{}
	}
	if err := input.Err(); err != nil {
		log.Println("读取错误：", err)

	}

	//用户离开
	leavingChannel <- user
	msg.Content = "user:`" + strconv.Itoa(user.ID) + "` has left"
	messageChannel <- msg

}

func sendMessage(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}

}

var (
	globalID int
	idLocker sync.Mutex
)

func GenUserID() int {
	idLocker.Lock()
	defer idLocker.Unlock()
	globalID++
	return globalID
}
