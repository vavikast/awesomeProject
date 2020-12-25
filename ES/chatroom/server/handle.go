package server

import (
	"awesomeProject/ES/chatroom/logic"
	"net/http"
)

var rootDir string

func RegisterHandler() {

	// 广播消息处理
	go logic.Broadcaster.Start()
	http.HandleFunc("/", homeHandleFunc)
	http.HandleFunc("/user_list", userListHandleFunc)
	http.HandleFunc("/ws", WebSocketHandleFunc)
}
