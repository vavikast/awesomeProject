package server

import (
	"awesomeProject/ES/chatroom/logic"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func homeHandleFunc(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles(rootDir + "/template/home.hmtl")
	if err != nil {
		fmt.Fprint(w, "模板解析错误")
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		fmt.Fprint(w, "模板执行错误")
		return
	}

}
func userListHandlefunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	userList := logic.Broadcaster.GetUserList()
	b, err := json.Marshal(userList)

	if err != nil {
		fmt.Fprint(w, `[]`)
	} else {
		fmt.Fprint(w, string(b))
	}

}
