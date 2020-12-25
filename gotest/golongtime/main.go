package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetTime() {
	now := time.Now()
	resp, err := http.Get("https://m.tianxiaxinyong.com/cooperation/crp-v2/index.html?txxychannel=eXQ5eWdDaVE1clNWTWt5RHc5Um1vdz09&txxysp=1")
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("打印信息错误")
	}
	duration := time.Since(now)
	fmt.Println(duration)

}
func main() {
	tick := time.Tick(time.Second * 1)
	exit := make(chan bool, 0)
	go func() {
		for {
			select {
			case <-tick:
				GetTime()
			case <-exit:
				return
			}
		}

	}()
	time.Sleep(12 * time.Hour)
	exit <- true

	fmt.Println("It's down")
	//readAll, err := ioutil.ReadAll(resp.Body)
	//if err != nil {bu
	//	fmt.Println("文件读取错误")
	//	return
	//}
	//fmt.Println(string(readAll))

}
