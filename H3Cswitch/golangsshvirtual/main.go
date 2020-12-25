package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

func main()  {
	sshConfig := &ssh.ClientConfig{
		Config:            ssh.Config{},
		User:              "root",
		Auth:              []ssh.AuthMethod{
			ssh.Password(`Itcen2531,.`),
		},
		HostKeyCallback:   ssh.InsecureIgnoreHostKey(),
	}
	connection, err := ssh.Dial("tcp", "192.168.14.137:22", sshConfig)
	if err != nil {
		log.Fatalln("建立连接错误",err)
	}
	session, err := connection.NewSession()
	defer  session.Close()
	if err != nil {
		log.Fatalln("建立会话错误",err)
	}
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	err = session.RequestPty("xterm", 80, 40, modes)
	if err != nil {
		log.Fatalln("创建pty出错",err)
	}
	pipe, _:= session.StdinPipe()
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	err = session.Shell()
	if err != nil {
		log.Fatalln("创建shell出错",err)
	}


	fmt.Fprintf(pipe,"\n%s\n%s\n%s\n","pwd","ls","exit")
	if err != nil {
		log.Fatalln("写入错误",err)
	}


	session.Wait()



}
