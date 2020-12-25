package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	_ "os"
	_ "regexp"
)

func main() {
	// 1.抓取主页
	//example  https://studygolang.com/topics?p=3

	//urlbase := "https://studygolang.com/topics?p="
	urlbase := "https://studygolang.com/topics/11598"
	resp, err := http.Get(urlbase)
	ErrPrint(err)
	defer resp.Body.Close()
	htmltxt, err := ioutil.ReadAll(resp.Body)
	pageContent := string(htmltxt)
	fmt.Println(pageContent)

	// 2.爬去go技术日报+爬取go相关公众号
	//设置正则匹配Go技术日报 <a href="/topics/11598" title="Go技术日报(2020-05-30)">Go技术日报(2020-05-30)</a>
	//re := regexp.MustCompile(`<a href="(/\w+/\d+)"\s+title="(\w+\p{Han}+\(.*?\))"`)
	//comments := re.FindAllStringSubmatch(pageContent, -1)
	//for _, comment := range comments {
	//	fmt.Println(comment[1], comment[2])
	//	//suburl := "https://studygolang.com/"+comment[1]
	//	baseurl := `https://studygolang.com/topics/`
	//	basefile:=`F:\goproject\src\awesomeProject\gocrawler\studygolang\`
	//	file, err := os.OpenFile(basefile+string(comment[2]), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//	ErrPrint(err)
	//	defer  file.Close()
	//	file.WriteString(baseurl+string(comment[1]))




		//WriteFile(file,basefile+string(comment[1]))
	//}
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