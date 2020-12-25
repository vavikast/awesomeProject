package main

import (
	"fmt"
	"os/exec"
	"time"
)

var x, y int

func NetWorkStatus() {
	cmd := exec.Command("ping", "m.tianxiaxinyong.com", "-n", "1", "-w", "1")
	now := time.Now()
	err := cmd.Run()
	duration := time.Since(now)
	fmt.Println(duration)
	y += 1
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	x += 1

}

func main() {
	tick := time.Tick(time.Second * 1)

	for {
		select {
		case <-tick:
			NetWorkStatus()
			fmt.Printf("总发送数据包: %v,成功接收数据包: %v,成功发送率是%0.2f%%\n", y, x, float64(x)/float64(y)*100)
		}

	}
}
