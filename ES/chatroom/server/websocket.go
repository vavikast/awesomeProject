package server

import (
	"awesomeProject/ES/chatroom/logic"
	"log"
	"net/http"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func WebSocketHandleFunc(w http.ResponseWriter, req *http.Request) {
	conn, err := websocket.Accept(w, req, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		log.Println("websocket accept error:", err)
		return
	}
	//1.新用户进来，构建用户实例
	token := req.FormValue("token")
	nickname := req.FormValue("nickname")
	if l := len(nickname); l < 2 || l > 20 {
		log.Println("nickname illegal:", nickname)
		wsjson.Write(req.Context(), conn, logic.NewErrorMessage("非法昵称，昵称长度： 2-20"))
		conn.Close(websocket.StatusUnsupportedData, "nickname illegal!")
		return
	}
	if !logic.Broadcaster.CanEnterRoom(nickname) {
		log.Println("昵称已存在：", nickname)
		wsjson.Write(req.Context(), conn, logic.NewErrorMessage("该昵称已存在！"))
		conn.Close(websocket.StatusUnsupportedData, "nickname exists")
		return
	}
	userHasToken := logic.NewUser(conn, token, nickname, req.RemoteAddr)

	//开启给用户发送的消息
	go userHasToken.SendMessage(req.Context())

	//给当前用户发送欢迎信息
	userHasToken.MessageChannel <- logic.NewWelcomeMessage(userHasToken)

	//避免token泄露
	tmpUser := *userHasToken
	user := &tmpUser
	user.Token = ""

	//给所有用户告知新用户对的到来
	msg := logic.NewWelcomeMessage(user)
	logic.Broadcaster.Broadcast(msg)

	//将用户加入广播器列表
	logic.Broadcaster.UserEntering(user)
	log.Println("user:", nickname, "joins chat")

	//接收用户消息
	err = user.ReceiveMessage(req.Context())

	//用户离开
	logic.Broadcaster.UserLeaving(user)
	msg = logic.NewUserLeaveMessage(user)
	logic.Broadcaster.Broadcast(msg)
	log.Println("user:", nickname, "leaves chat")
	if err == nil {
		conn.Close(websocket.StatusNormalClosure, "")
	} else {
		log.Println("read from client error:", err)
		conn.Close(websocket.StatusInternalError, "Read from client error")
	}

}
