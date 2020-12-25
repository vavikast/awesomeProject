package main

import (
	"fmt"
	"net/http"
	"strings"

	//"regexp"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	// 1.抓取主页
	//example  https://studygolang.com/topics?p=3

	urlbase := "https://studygolang.com/topics/11598"
	resp, err := http.Get(urlbase)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	doc,_:= goquery.NewDocumentFromReader(resp.Body)
	htm,_ := doc.Find(".markdown-body").Html()


	//解码
	htm = decode(htm)

	//再次查询
	doc,_= goquery.NewDocumentFromReader(strings.NewReader(htm))
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		fmt.Printf("%v: %s\n",s.Text(),link)
	})


}

func decode(s string) string {
	s = strings.Replace(s,"&lt;","<",-1)
	s = strings.Replace(s,"&gt;",">",-1)
	s = strings.Replace(s,"&#34;","\"",-1)
	s = strings.Replace(s,"&#39;",";",-1)
	return  s
}
