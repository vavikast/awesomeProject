package main

import (
	"fmt"
	"strings"
	"time"
)

func main()  {
	var  a string
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("请输入字符串")
		return
	}
	length := len(a)
	//条件判断确定mac长度，包括三种格式的长度，不满足则打印格式错误
	if length == 12 || length == 15 ||length == 14 || length == 17{
		//如果是：格式mac则转换为标准格式
		//如果是-格式mac则转换为标准格式
		//如果是标准格式，根据选择分别转换成响应格式
		if strings.Contains(a,":") {
			a = strings.ToLower(a)
			a = strings.ReplaceAll(a,":","")
			fmt.Println(a)

		} else 	if strings.Contains(a,"-") {
			a = strings.ToLower(a)
			a = strings.ReplaceAll(a,"-","")
			fmt.Println(a)

		} else {
			fmt.Println("请输入正确的数字，1表示将标准mac转换为带:格式，2表示将标准mac转换为-格式:")
			a = strings.ToLower(a)
			var b int
			_, err := fmt.Scan(&b)
			if err != nil {
				fmt.Println("请输入正确的数字，1表示将标准mac转换为带:格式，2表示将标准mac转换为-格式")
				return
			}
			if b==1 {
				c := []rune(a)
				j := []rune{}
				d := c[:2]
				e := c[2:4]
				f := c[4:6]
				g := c[6:8]
				h := c[8:10]
				i := c[10:]
				j = append(append(append(append(append(append(append(append(append(append(append(j,d...),':'),e...),':'),f...),':'),g...),':'),h...),':'),i...)
				fmt.Println(string(j))

			} else if  b== 2 {
				c := []rune(a)
				g := []rune{}
				d := c[:4]
				e := c[4:8]
				f := c[8:]
				g = append(append(append(append(append(g,d...),'-'),e...),'-'),f...)
				fmt.Println(string(g))
			}else {
				fmt.Println("输入错误，请输入正确的数字，1:表示将标准mac转换为带:格式，2:表示将标准mac转换为-格式")
			}
		}
		time.Sleep(10*time.Second)
	} else {
		fmt.Println("格式错误")
	}


}
