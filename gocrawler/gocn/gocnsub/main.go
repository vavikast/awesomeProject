package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	// 1.抓取主页
	//example  https://gocn.vip/topics?page=

	urlbase := "https://gocn.vip/topics/10475"
	resp, err := http.Get(urlbase)
	ErrPrint(err)
	defer resp.Body.Close()
	htmltxt, err := ioutil.ReadAll(resp.Body)
	pageContent := string(htmltxt)
	//fmt.Println(pageContent)
	// 2.爬去GoCN每日新闻
	re := regexp.MustCompile(`<li>(.*)<a href="(.*)"\s+rel`)
	comments := re.FindAllStringSubmatch(pageContent, -1)
	for _, comment := range comments {
		if !strings.Contains(comment[1],"订阅新闻") && !strings.Contains(comment[1],"招聘专区"){
			fmt.Println(comment[1],comment[2])
		}

	}
	////comments := re.FindAllString(pageContent, -1)
	//fmt.Println(comments)
	// 3.拿到里面的标题和连接
	//4.分别存到对应的日期的文件下面
}

//func WriteFile(file *os.File, url string) (err error) {
//	resp, err := http.Get(url)
//	ErrPrint(err)
//	defer resp.Body.Close()
//	htmltxt, err := ioutil.ReadAll(resp.Body)
//	pageContent := string(htmltxt)
//
//}
func ErrPrint(err error)  {
	if err != nil {
		panic(err)
	}
}
