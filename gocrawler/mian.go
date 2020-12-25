package main

import (
	"fmt"
	"net/http"
	"time"
)

func main()  {
	var lastRequestTime time.Time

	//最大请求数量
	maximumNumberOfRequests := 5

	//每5秒1页
	pageDelay := 5* time.Second

	for i:=0;i<maximumNumberOfRequests;i++{
		elapsedTime := time.Now().Sub(lastRequestTime)
		fmt.Printf("Elapsed Time: %.2f (s)\n",elapsedTime.Seconds())
		if elapsedTime < pageDelay{
			var timeDiff time.Duration = pageDelay- elapsedTime
			fmt.Printf("Sleeping for %.2f ()\n",timeDiff.Seconds())
			time.Sleep(pageDelay-elapsedTime)
		}
		fmt.Println("Get studygolang ")
		_, err := http.Get("https://studygolang.com/")
		if err != nil {
			panic(err)

		}
		lastRequestTime = time.Now()
	}



}