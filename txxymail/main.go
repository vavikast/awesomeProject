package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	//"strconv"
)

//func Index(w http.ResponseWriter, r *http.Request) {
//	// fmt.Fprintf(w, "Hello world, this is my first page!")
//
//	html := `<!DOCTYPE html>
//<html lang="en">
//
//<head>
//    <meta charset="UTF-8">
//    <meta name="viewport" content="width=device-width, initial-scale=1.0">
//    <title>Form</title>
//    <style>
//        body{
//            display: flex;
//            justify-content: center;
//            align-items: center;
//        }
//        input,
//        textarea {
//            display: block;
//            margin-top: 16px;
//            font-size: 18px;
//        }
//        textarea {
//            width: 500px;
//            height: 100px;
//        }
//        input {
//            width: 500px;
//            height: 40px;
//            line-height: 40px;
//        }
//    </style>
//</head>
//
//<body>
//<form method="get" action="http://www.baidu.com">
//    <input type="text" name="name" placeholder="name">
//    <input type="text" name="email" placeholder="email">
//    <input type="text" name="subject" placeholder="subject">
//    <textarea name="message" placeholder="message"></textarea>
//    <input type="submit" value="send">
//</form>
//</body>
//
//</html>`
//	fmt.Fprintln(w, html)
//}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求的方法


	if r.Method == "GET" {
			tem, err := template.ParseFiles(`F:\goproject\src\awesomeProject\txxymail\template\index.html`)
		if err != nil {
			fmt.Println("读取文件错误，err:",err)
			return
		}
		tem.Execute(w,nil)
	} else {
		err := r.ParseForm()   // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
		if err != nil {
			// handle error http.Error() for example
			log.Fatal("ParseForm: ", err)
		}
		// 请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("mailto:", r.Form["mailto"])

		fmt.Println("subject:", r.Form["subject"])
		fmt.Println("body:", r.Form["body"])

	}
}
func login(w http.ResponseWriter, r *http.Request) {
	var mailto string
	var subject string
	var message string
	var username,password string
	fmt.Println("method:", r.Method) // 获取请求的方
	if r.Method == "GET" {
		err := r.ParseForm()   // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
		if err != nil {
			// handle error http.Error() for example
			log.Fatal("ParseForm: ", err)
		}
		// 请求的是登录数据，那么执行登录的逻辑判断

		mailto = fmt.Sprintln(r.Form["mailto"][0])
		subject =fmt.Sprintln(r.Form["subject"][0])
		fmt.Println("subject:", r.Form["subject"])
		fmt.Println("message:", r.Form["message"])
		message =fmt.Sprintln(r.Form["message"][0])
		tem, err := template.ParseFiles(`F:\goproject\src\awesomeProject\txxymail\template\login.html`)
		if err != nil {
			fmt.Println("读取文件错误，err:",err)
			return
		}
		data := map[string]interface{}{"Mailto":mailto,"Subject":subject,"Message":message}
		tem.Execute(w,data)
	}

	if r.Method == "POST" {
		err := r.ParseForm()   // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
		if err != nil {
			// handle error http.Error() for example
			log.Fatal("ParseForm: ", err)
		}
		s,_ := ioutil.ReadAll(r.Body)
		fmt.Println(s)
		username = r.PostFormValue("username")
		password = r.PostFormValue("password")
		mailto = r.PostFormValue("mailto")
		subject = r.PostFormValue("subject")
		message = r.PostFormValue("message")
		//username = fmt.Sprintln(r.Form["username"])
		//password = fmt.Sprintln(r.Form["password"])
		//password = fmt.Sprintln(r.Form["mailto"])
		//password = fmt.Sprintln(r.Form["subject"])
		//password = fmt.Sprintln(r.Form["message"])
		fmt.Println(username)
		fmt.Println(password)
		fmt.Println(mailto)
		fmt.Println(subject)
		fmt.Println(message)
		err = Checkmail(username, password, mailto, subject, message)
		if err != nil {
			fmt.Fprintf(w, "邮件发送成功!")
		} else {
			fmt.Fprintf(w, "邮件发送失败!")
		}


	}

}

func Checkmail(username,password,mailto,subject,message string)  error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	port := 465

	//port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	m.SetHeader("From","天下信用测试邮箱" + "<" + username + ">")  //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailto)  //发送给多个用户
	m.SetHeader("Subject", subject)  //设置邮件主题
	m.SetBody("text/html", message)     //设置邮件正文

	d := gomail.NewDialer("smtp.qiye.aliyun.com", port, username, password)

	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println("发送成功")
	}
	return err
}



func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)

	//dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	//fmt.Println(dir)
	// 监听本机的8030端口
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error: ", err)
	}




}
