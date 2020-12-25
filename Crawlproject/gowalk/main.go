package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"strings"
)

func main() {
	//声明两个文本域控件
	var inTE, outTE *walk.TextEdit

	//配置主窗口，运行起来
	MainWindow{
		//窗口标题
		Title: "尖叫的demo",
		//可拉伸的最小尺寸
		MinSize: Size{600, 400},

		//主布局： 垂直布局
		Layout: VBox{},

		//窗口中的所有控件
		Children: []Widget{

			//水平分割器（水平小布局）
			HSplitter{

				//局部水平排列的控件们
				Children: []Widget{
					//文本输入框
					TextEdit{
						//绑定到inTE变量
						AssignTo: &inTE},

					//文本输出框
					TextEdit{
						AssignTo: &outTE,
						//只读的文本框
						ReadOnly: true},
				},
			},

			//普通按钮
			PushButton{

				//按钮文本
				Text: "动动我试试",

				//响应函数
				OnClicked: func() {
					inputStr := inTE.Text()
					outputStr := strings.ToUpper(inputStr)
					outTE.SetText(outputStr)
				},
			},
		},
	}.Run()

}
