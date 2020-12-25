package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"strings"
	"sync"
)


//获取账号和密码的对应关系
type HostPassword struct {
	Host string
	Username string
	Password string
}
var (
	a,b string  //临时存储变量
	commands = []string{}  //执行命令组
	hp []HostPassword  //保存账号和密码
	wg sync.WaitGroup	 //执行goroutine

)
func main()  {
	//1. 选择交换机
	//2. 输入要执行命令
	//3. 建立会话连接
	//4. 新建session，并执行命令

	//1. 选择操作交换机
	// 1.1 输入要执行交换机
	fmt.Println("请输入计划执行命令的交换机账号和密码,账号密码直接使用|分割，多个账号密码之间使用;分割,例如admin;192.168.56.10;h3csw1|admin;192.168.56.11;h3csw2")
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Fatal("输入错误:",err)
	}
	fmt.Println("请输入要执行的命令行，以；号间隔")
	//1.1.1切割交换机命令
	switchgroups := strings.Split(a, "|")
	length := len(switchgroups)
	hp = make([]HostPassword,length)
	for i,singleswitch := range switchgroups{
		hp[i]=HostPassword{}
		switchsplit := strings.Split(singleswitch, ";")
		hp[i].Username=switchsplit[0]
		hp[i].Host=switchsplit[1]
		hp[i].Password=switchsplit[2]
	}


	// 1.2 输入要执行命令
	input := bufio.NewReader(os.Stdin)
	b, err := input.ReadString('\n')
	if err != nil {
		log.Fatal("输入错误",err)
	}
	commands = strings.Split(b, ";")


	//2. 执行交换机操作
	err = SshSwitch(hp)
	if err != nil {
		log.Fatalln(err)
	}

	// 同步等待
	wg.Wait()
}


//建立ssh连接
func SshSwitch(hostpasswords []HostPassword) (error){
	//循环获取hostpasswords的账号和密码
	for i,_ := range hp{
		//添加同步组，下面会执行goroutin
		wg.Add(1)
		config := &ssh.ClientConfig{
			Config:            ssh.Config{
				Ciphers:        []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
			}, //添加了很多加密方式，为了应对不同的密码规则
			User:              hp[i].Username,
			Auth:              []ssh.AuthMethod{
				ssh.Password(hp[i].Password),
			},
			HostKeyCallback:   ssh.InsecureIgnoreHostKey(), //此处相当于执行nil，但是并不安全
		}
		client, err := ssh.Dial("tcp",hp[i].Host+":22", config)
		if err != nil {
			log.Fatalln("建立ssh连接错误:",err)
			return err
		}
		//执行goroutine，但是没有返回错误。
		go HandleSession(client, commands,&wg)

	}
	return  nil
}

//建立session，执行命令。
func HandleSession(client *ssh.Client,commands []string,wg *sync.WaitGroup) error {
	//建立session
	session, err := client.NewSession()
	if err != nil {
		log.Fatalln("创建session出错",err)
		return  err
	}
	//延迟关闭session
	defer  session.Close()

	//设置terminalmodes的方式
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	//建立伪终端
	err = session.RequestPty("xterm",80,40,modes)
	if err != nil {
		log.Fatal("创建requestpty出错",err)
		return  err
	}
	//设置session的标准输入是stdin
	stdin, err := session.StdinPipe()
	if err != nil {
		log.Fatal("输入错误",err)
		return  err
	}
	//设置session的标准输出和错误输出分别是os.stdout,os,stderr.就是输出到后台
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	err = session.Shell()
	if err != nil {
		log.Fatal("创建shell出错",err)
		return  err
	}
	//将命令依次执行
	for _, cmd := range commands {
		fmt.Println(cmd)
		_, err = fmt.Fprintf(stdin, "%s\n", cmd)
		if err != nil {
			log.Fatal("写入stdin出错",err)
			return  err
		}
	}

	//执行等待
	err = session.Wait()
	if err != nil {
		log.Fatal("等待session出错",err)
		return  err
	}
	//减少同步组的次数
	wg.Done()
	return  nil
}

