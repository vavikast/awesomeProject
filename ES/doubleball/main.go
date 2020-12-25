package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"math/rand"
	"sort"
	"time"
)

var (
	redball    = make(map[int]int) //红球
	bullball   = make(map[int]int) // 篮球
	doubleabll = make([]int, 0)    // 双色球
)

func main() {
	//设置text的编辑模式
	var usernameTE, passwordTE, dateTE *walk.TextEdit

	//设置主窗口
	MainWindow{
		Title:   "天天双色球，财富到我手",
		MinSize: Size{400, 50},
		Size:    Size{400, 50}, //这里才是显示窗口大小的关键
		Layout:  VBox{},        // 竖直布局模式
		Children: []Widget{
			Composite{
				Layout: HBox{}, //水平布局模式
				Children: []Widget{
					GroupBox{
						//Title:  "选择号码",中文显示不全，所以取消了
						Layout: VBox{},
						Children: []Widget{
							Label{Text: "红球区:"},
							TextEdit{
								//MaxSize: Size{30,10},
								AssignTo: &usernameTE,
								ReadOnly: true,
							},
							Label{
								//MaxSize: Size{100, 40},
								Text: "篮球区:",
							},
							TextEdit{
								//MaxSize: Size{30,10},
								AssignTo: &passwordTE,
								ReadOnly: true,
							},
						},
					},
					GroupBox{
						//Title:  "开奖日期", 中文显示不全，所以取消了
						Layout: VBox{},
						Children: []Widget{
							Label{Text: "开奖日期:"},
							TextEdit{
								//MaxSize: Size{30,10},
								AssignTo: &dateTE,
								ReadOnly: true,
							},
						},
					},
				},
			},
			// 按钮
			PushButton{
				Text:       "幸运之门",
				Persistent: true,
				//MinSize: Size{30, 10},
				OnClicked: func() {
					//获取双色球的随机数
					doubleabll = Randball()
					a := fmt.Sprint(doubleabll[:6])
					b := fmt.Sprint(doubleabll[6:])
					usernameTE.SetText(a)
					passwordTE.SetText(b)
					dateTE.SetText(GetDate())
					//下面三项是关键，恢复默认值，要不然很很尴尬。
					doubleabll = []int{}
					redball = map[int]int{}
					bullball = map[int]int{}
				},
			},
		},
	}.Run()

}

//获取双色球随机数
func Randball() []int {
	//种子
	rand.Seed(time.Now().Unix())

	//获取6个红球
	for j := 0; j < 6; j++ {
		//左闭右开[0 33)
		intn := rand.Intn(33)
		//通过map方式，将红球存成[2]3
		//为了避免冲突，如果已经在map中存在此数值，则顺找下一个
		if _, ok := redball[intn]; ok {
			j--
			fmt.Println(redball)
		} else {
			fmt.Println(j)
			redball[intn] = intn + 1
			doubleabll = append(doubleabll, redball[intn])
		}

	}
	//对切片进行排序，是对原始切片排序
	sort.Ints(doubleabll)
	for i := 0; i < 1; i++ {
		intn := rand.Intn(16)
		if _, ok := bullball[intn]; ok {
			i--
		} else {
			bullball[intn] = intn + 1
			doubleabll = append(doubleabll, bullball[intn])
		}

	}

	return doubleabll
}

//获取开奖日期，这里是从今天往后，双色器第一次开奖时间。 开奖时间为周二，周四，周日
func GetDate() string {
	now := time.Now()
	//遇到周二，周四，周日返回
	for {
		weekday := GetWeekday(now)
		if weekday == "Tuesday" || weekday == "Thursday" || weekday == "Sunday" {
			weekname := ""
			if weekday == "Tuesday" {
				weekname = "星期二"
			}
			if weekday == "Thursday" {
				weekname = "星期四"
			}
			if weekday == "Sunday" {
				weekname = "星期天"
			}
			return fmt.Sprint(now.Format("2006-01-02")) + " " + weekname
		} else {
			now = now.Add(24 * time.Hour)
		}
	}
}

//将时间转换为字符串模式
func GetWeekday(t time.Time) string {
	Weekday := fmt.Sprint(t.Weekday())
	return Weekday
}
