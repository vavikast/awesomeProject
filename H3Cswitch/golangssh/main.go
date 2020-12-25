package main

import (
	"bytes"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

const (
	username = "huawei"
	password = "huawei"
	connecthost = "192.168.56.10:22"
	cmds = "show clock;show env power;exit"
)
func connect(user,password,connecthost,key string,cipherList []string) (*ssh.Session,error)  {
	var (
		auth []ssh.AuthMethod
		clientConfig *ssh.ClientConfig
		config ssh.Config
		session *ssh.Session
		client       *ssh.Client
		err error
	)
	auth = make([]ssh.AuthMethod,0)
	if key == ""{
		auth = append(auth,ssh.Password(password))
	} else {
		pemBytes, err := ioutil.ReadFile(key)
		if err != nil {
			return nil,err
		}
		var signer ssh.Signer
		if password == ""{
			signer,err = ssh.ParsePrivateKey(pemBytes)
		} else {
			signer,err = ssh.ParsePrivateKeyWithPassphrase(pemBytes,[]byte(password))
		}
		if err != nil {
			return nil,err
		}
		auth = append(auth,ssh.PublicKeys(signer))
	}

	if len(cipherList) == 0{
		config = ssh.Config{
			Ciphers:        []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
		}
	} else {
		config = ssh.Config{
			Ciphers:  cipherList,
		}
	}
	clientConfig = &ssh.ClientConfig{
		Config:            config,
		User:              user,
		Auth:              auth,
		HostKeyCallback:   ssh.InsecureIgnoreHostKey(),
	}

	client, err = ssh.Dial("tcp", connecthost, clientConfig)
	if err != nil {
		return nil,err
	}
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	err = session.RequestPty("xterm", 80, 40, modes)
	if err != nil {
		return nil,err
	}
	return  session,err

}

func main()  {
	var cipherList []string
	session, err := connect(username, password, connecthost, "", cipherList)
	if err != nil {
		log.Fatal(err)
	}
	defer  session.Close()

	cmdlist := strings.Split(cmds, ";")
	stdinBuf, err := session.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	var outbt,errbt bytes.Buffer
	session.Stdout = &outbt
	session.Stderr = &errbt
	for _,c := range cmdlist{
		c = c+"\n"
		stdinBuf.Write([]byte(c))
	}
	time.Sleep(1*time.Minute)
	session.Wait()


}