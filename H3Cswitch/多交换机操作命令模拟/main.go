package main

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	b string  //临时存储变量
	username = "h3c" //用户名
	password1 = "h3csw1" //密码1
	password2 = "h3csw2"
	HostPasswods = make(map[string]string) //存储主机和密码对应关系
	commands = []string{}  //执行命令组
	wg sync.WaitGroup	 //执行goroutine
	switchTurn = false //交换机mac地址添加开关
	AcTurnOn = false  //AC mac地址添加开关
	AcTurnOff = false //AC mac地址删除开关
)
func main()  {
	//1. 选择交换机
	//2. 输入要执行命令
	//3. 建立连接
	//4. 新建session，并执行命令

	//1. 选择操作交换机
	// 输入要进行命令的交换机
	var a string
	fmt.Println("请输入要操作的交换机，以逗号方式输入，例如1,2。 核心交换机代表1,办公交换A表示2，办公交换机B表示3,办公交换机C表示4,服务器交换机表示5，AC表示6。快捷操作必备: 交换机添加mac地址输入7,AC添加mac输入8,AC删除mac输入9")
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Fatal("输入错误:",err)
	}

	// 获取要执行交换机的编号
	temp := strings.Split(a, ",")
	for _,i := range  temp{
		num, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal("转换数字出错",err)
		}
		//根据序号,添加匹配主机
		err = Addnum(num)
		if err != nil {
			log.Fatal(err)
		}
	}
	commands = GetCommand()
	//2. 执行交换机操作
	err = SshSwitch(HostPasswods)
	if err != nil {
		log.Fatalln(err)
	}

	// 同步等待
	wg.Wait()
}

//生成主机-密码的map
func Addnum(num int) (error) {
	if num == 1 {
		HostPasswods["192.168.56.10:22"]=password1
	} else   if num == 1 {
		HostPasswods["192.168.56.11:22"]=password2
	} else   if num == 2 {
		HostPasswods["192.168.56.11:22"]=password2
	} else   if num == 3 {
		HostPasswods["192.168.56.11:22"]=password2
	} else   if num == 4 {
		HostPasswods["192.168.56.11:22"]=password2
	} else   if num == 5 {
		HostPasswods["192.168.56.11:22"]=password2
	} else   if num == 6 {
		HostPasswods["192.168.56.11:22"]=password2
	} else if num == 7{
		HostPasswods["192.168.56.10:22"]=password1
		HostPasswods["192.168.56.11:22"]=password2
		switchTurn = true
	} else if num == 8 {
		HostPasswods["192.168.56.10:22"]=password1
		AcTurnOn = true
	} else if  num == 9 {
		HostPasswods["192.168.56.10:22"]=password1
		AcTurnOff = true
	} else	{
		err := errors.New("没有需求数据，输入错误")
		return err
	}
	return nil
}

//获取commands命令
func  GetCommand() []string  {
	if switchTurn {
		var c string
		var d int
		fmt.Println("请输入mac地址:")
		_, err := fmt.Scanln(&c)
		if err != nil {
			log.Fatal("输入mac错误:",err)
		}
		fmt.Println("输入vlan id:")
		_, err = fmt.Scanln(&d)
		if err != nil {
			log.Fatal("输入vlan错误:",err)
		}
		//转换mac地址格式
		temp := SwitchMac(c,2)
		//执行命令
		b = `system-view;undo mac-vlan mac-address `+temp+`;mac-vlan mac-address `+temp+`  vlan `+strconv.Itoa(d)+`;save force;quit`
		//将命令添加切片
		commands = strings.Split(b, ";")
		fmt.Println(commands)
	} else  if AcTurnOn{
		var c string
		fmt.Println("请输入mac地址:")
		_, err := fmt.Scanln(&c)
		if err != nil {
			log.Fatal("输入mac错误:",err)
		}
		temp := SwitchMac(c,0)
		// 操作命令
		b = `system-view;local-user `+temp+` class network;password simple  `+temp+`;service-type lan-access;quit;save force;quit`
		commands = strings.Split(b, ";")
		fmt.Println(commands)

	} else if AcTurnOff {
		var c string
		fmt.Println("请输入要删除mac地址")
		_, err := fmt.Scanln(&c)
		if err != nil {
			log.Fatal("输入mac错误",err)
		}
		temp := SwitchMac(c,0)
		// 操作命令
		b = `system-view;undo local-user `+temp+` class network;quit;save force`
		commands = strings.Split(b, ";")
		fmt.Println(commands)
	} else {
		//获取执行交换机的命令，之所以这样主要是要输出的操作命令有空格，scan不能获取
		fmt.Println("请输入要执行的命令行，以；号间隔")
		input := bufio.NewReader(os.Stdin)
		b, err := input.ReadString('\n')
		if err != nil {
			log.Fatal("输入错误",err)
		}
		commands = strings.Split(b, ";")
	}
	return commands
}

//建立ssh连接
func SshSwitch(hostpasswords map[string]string) (error){
	//循环获取hostpasswords的账号和密码
	for host,password := range HostPasswods{
		//添加同步组，下面会执行goroutin
		wg.Add(1)
		config := &ssh.ClientConfig{
			Config:            ssh.Config{
				Ciphers:        []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
			}, //添加了很多加密方式，为了应对不同的密码规则
			User:              username,
			Auth:              []ssh.AuthMethod{
				ssh.Password(password),
			},
			HostKeyCallback:   ssh.InsecureIgnoreHostKey(), //此处相当于执行nil，但是并不安全
		}
		client, err := ssh.Dial("tcp",host, config)
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

//mac地址转换，如果ant=0表示输出去除格式的mac，如果ant=1表示以:的mac，如果ant=2表示以-分割的mac
func SwitchMac(a string,ant int) (mac string) {
	length := len(a)
	//条件判断确定mac长度，包括三种格式的长度，不满足则打印格式错误
	if length == 12 || length == 14 || length == 17 {
		//如果是：格式mac则转换为标准格式
		//如果是-格式mac则转换为标准格式
		//如果是标准格式，根据选择分别转换成响应格式
		if strings.Contains(a, ":") {
			a = strings.ToLower(a)
			a = strings.ReplaceAll(a, ":", "")
			fmt.Println(a)
		}
		if strings.Contains(a, "-") {
			a = strings.ToLower(a)
			a = strings.ReplaceAll(a, "-", "")
		}
		a = strings.ToLower(a)
		if ant == 1 {
			c := []rune(a)
			j := []rune{}
			d := c[:2]
			e := c[2:4]
			f := c[4:6]
			g := c[6:8]
			h := c[8:10]
			i := c[10:]
			j = append(append(append(append(append(append(append(append(append(append(append(j, d...), ':'), e...), ':'), f...), ':'), g...), ':'), h...), ':'), i...)
			a = string(j)
			return a
		} else if ant == 2 {
			c := []rune(a)
			g := []rune{}
			d := c[:4]
			e := c[4:8]
			f := c[8:]
			g = append(append(append(append(append(g, d...), '-'), e...), '-'), f...)
			a = string(g)
			return a
		} else {
			return a
		}
	} else {
		panic("格式错误")
	}
}
