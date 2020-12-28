package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//钉钉自定义机器人： https://ding-doc.dingtalk.com/document#/org-dev-guide/custom-robot/title-72m-8ag-pqw
//{
//	"msgtype": "text",
//	"text": {
//		"content": "我就是我, @150XXXXXXXX 是不一样的烟火"
//},
//	"at": {
//		"atMobiles": [
//		"150XXXXXXXX"
//],
//	"isAtAll": false
//}
//}
//定义匹配json结构体格式
type Text struct {
	Content string `json:"content"`
}

type Webhookjson struct {
	Msgtype string `json:"msgtype"`
	Text    `json:"text"`
}

func main() {
	text := "我就是我, 是不一样的烟火"
	SendWebhook(text)

}

//转换为json格式。
func ToJson(text string) (marshal []byte, err error) {
	newjson := Webhookjson{
		Msgtype: "text",
	}
	newjson.Content = text
	marshal, err = json.Marshal(newjson)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	return marshal, nil
}

//发送webhook信息
func SendWebhook(text string) error {
	marshaljson, err := ToJson(text)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//钉钉webhook
	url := "https://oapi.dingtalk.com/robot/send?access_token=51345145d106753486bd71614bf881283f91e2124535276b257f99327e41dc87"
	//钉钉格式
	contentType := "application/json"
	resp, err := http.Post(url, contentType, bytes.NewBuffer(marshaljson))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))
	return nil
}
