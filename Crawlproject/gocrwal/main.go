package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	//邮箱
	reQQmail = `(\d+)@qq.com`
	reEmail  = `\w+@\w+\.\w+(\.\w+)?`
	//超链接
	//<a href="http://news.baidu.com/ns?cl=2&rn=20&tn=news&word=%C1%F4%CF%C2%D3%CA%CF%E4%20%B5%BA%B9%FA"
	reLinkBad = `<a[\s\S]*?href="(https?://[\s\S]+？)"`
	reLink    = `href="(https?://[\s\S]+?)"`

	//手机号
	//13xxx xxx xxx
	rePhone = `1[2345789]\d\s?\d{4}\s?\d{4}`

	//s身份证号码
	reIdcard = `[123456]\d{5}((19\d{2})|(20[01]\d))((0[1-9])|(1[012]))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dX]`

	//图片链接
	//"http://img2.imgtn.bdimg.com/it/u=2403021088,4222830812&fm=26&gp=0.jpg"
	reImg = `"(https?://[^"]+?(\.((jpg)|(jpeg)|(png)|(gif)|(bmp)|(svg)|(swf)|(ico))))"`
)

//预处理错误
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}

}

//获取页面html方法封装。
func GetpageStr(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleError(err, "http,Get url")
	defer resp.Body.Close()
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil,Readall")
	pageStr = string(pageBytes)
	return pageStr

}

//爬邮箱
