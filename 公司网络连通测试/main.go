package main

import (
	"fmt"
	"github.com/go-gomail/gomail"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func main() {
	//获取主机列表
	list := GetNetworkList()
	for {
		NetWorkStatus(list)
		time.Sleep(5 * time.Minute)
	}

}

//网络侦测
func NetWorkStatus(networklist map[string]string) {
	for networkname, ipaddr := range networklist {
		fmt.Println(networkname, ipaddr)
		out, _ := exec.Command("ping", ipaddr, "-c", "5", "-i", "0", "-W", "1").Output()
		fmt.Println(string(out))
		if strings.Contains(string(out), "100% packet loss") {
			fmt.Println("network down")
			Sendmail(networkname, ipaddr)
		} else {
			fmt.Println("IT'S ALIVEEE")
		}
	}
}

//发送邮件
func Sendmail(networkname, ipaddr string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "wangw02@txxy.com")
	m.SetHeader("To", "wangw02@txxy.com")
	m.SetHeader("Subject", "公司网络连接故障，请查看!")
	now := time.Now()
	nowtime := now.Format("2006-01-02 15-04-05")
	htmlbody := nowtime + "  公司网络连接故障 " + networkname + ipaddr
	m.SetBody("text/html", htmlbody)

	d := gomail.NewDialer("smtp.qiye.aliyun.com", 25, "wangw02@txxy.com", "xxx")
	fmt.Println("发送邮件")
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}

//获取发送邮件发送列表
func GetNetworkList() map[string]string {
	data, err := ioutil.ReadFile(`F:\goproject\src\awesomeProject\公司网络连接测试\pattern\network.txt`)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	//PPPoE ethernet0/3_pppoe从服务器获得IP地址100.64.18.19\n\u0000
	r, _ := regexp.Compile(`PPPoE (?P<network>ethernet0/[1-9])_pppoe从服务器获得IP地址(?P<ipaddr>[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3})`)
	//networkipaddr
	networkipaddr := r.FindAllStringSubmatch(string(data), -1)

	ethaddr := make(map[string]string)
	//添加两个固定IP地址
	ethaddr["移动专线-ethernet0/4"] = "183.xx.158.xxx"
	ethaddr["电信ICP-ethernet0/7"] = "183.xx	.131.xxx"
	for _, onenetworkaddr := range networkipaddr {
		// 对每一行生成一个map
		ethaddr["电信ADSL"+"-"+onenetworkaddr[1]] = onenetworkaddr[2]
	}
	return ethaddr
}
