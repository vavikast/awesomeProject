package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var urlbase1 = "https://gocn.vip/topics?page="
var baseurl2 = "https://gocn.vip"
var basefile =`F:\goproject\src\awesomeProject\gocrawler\gocn\`
var wg sync.WaitGroup

func main() {
	// 1.抓取主页
	//example  https://gocn.vip/topics?page=
	for i := 1;i<=60;i++{
		wg.Add(1)
		url1:=urlbase1+strconv.Itoa(i)
		go mainGoDownload(&wg,url1)
	}
	wg.Wait()
	// 3.拿到里面的标题和连接
	//4.分别存到对应的日期的文件下面
}

func WriteFile(file *os.File, url string) (err error) {
	s := GoDownload(url)

	// 2.爬去GoCN每日新闻
	re := regexp.MustCompile(`<li>(.*)< a href="(.*)"\s+rel`)
	comments := re.FindAllStringSubmatch(s, -1)
	for _, comment := range comments {
		if !strings.Contains(comment[1],"订阅新闻") && !strings.Contains(comment[1],"招聘专区"){
			fmt.Println(comment[1],comment[2])
			surl := comment[1]+comment[2]+"\n"
			file.WriteString(surl)
		}
	}

	return  err
}
func ErrPrint(err error)  {
	if err != nil {
		panic(err)
	}
}

func GoDownload(url string) (s string) {
	resp, err := http.Get(url)
	ErrPrint(err)
	defer resp.Body.Close()
	htmltxt, err := ioutil.ReadAll(resp.Body)
	ErrPrint(err)
	s = string(htmltxt)
	return  s
}
func mainGoDownload(wg *sync.WaitGroup,url string)  (err error){
	pageContent := GoDownload(url)
	// 2.爬去GoCN每日新闻
	//设置正则匹配GoCN每日新闻 <a title="GoCN 每日新闻 (2020-05-30)" href=" ">
	re := regexp.MustCompile(`<a title="\w+\s?\p{Han}+\s?\((.*)\)"\s+href="(.*)"`)
	comments := re.FindAllStringSubmatch(pageContent, -1)
	for _, comment := range comments {
		fmt.Println(comment[1],comment[2])
		file, err := os.OpenFile(basefile+comment[1], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		ErrPrint(err)
		defer  file.Close()
		surl := baseurl2+string(comment[2])
		WriteFile(file,surl)

	}
	wg.Done()
	return err
}