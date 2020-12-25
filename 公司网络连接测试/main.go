package main

import (
	"fmt"
	"github.com/go-gomail/gomail"
	"os/exec"
	"strings"
	"time"
)

func main() {
	ip := "183.62.131.138"
	for {
		NetWorkStatus(ip)
		time.Sleep(5 * time.Minute)
	}

}

//网络侦测
func NetWorkStatus(ip string) {
	out, _ := exec.Command("ping", ip, "-c", "5", "-i", "0", "-W", "1").Output()
	fmt.Println(string(out))
	if strings.Contains(string(out), "100% packet loss") {
		fmt.Println("network down")
		Sendmail(ip)
	} else {
		fmt.Println("IT'S ALIVEEE")
	}

}

//发送邮件
func Sendmail(ip string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "wangw02@txxy.com")
	m.SetHeader("To", "wangw02@txxy.com")
	m.SetHeader("Subject", "公司网络连接故障，请查看!")
	now := time.Now()
	nowtime := now.Format("2006-01-02 15-04-05")
	htmlbody := nowtime + "  公司网络连接故障 " + ip
	m.SetBody("text/html", htmlbody)

	d := gomail.NewDialer("smtp.qiye.aliyun.com", 25, "wangw02@txxy.com", "xxxx")
	fmt.Println("发送邮件")
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
