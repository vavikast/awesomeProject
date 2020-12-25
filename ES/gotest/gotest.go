package main

import (
	"awesomeProject/ES/kubernetes/staging/src/k8s.io/apimachinery/pkg/util/rand"
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"sort"
	"time"
)

var (
	redball    = make(map[int]int)
	doubleabll = make([]int, 0)
)

func main() {

	var usernameTE, passwordTE, dateTE *walk.TextEdit

	MainWindow{
		Title:   "天天双色器，财富到我手",
		MinSize: Size{400, 50},
		Size:    Size{400, 50},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout: HBox{},
				Children: []Widget{
					GroupBox{
						//Title:  "选择号码",
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
						//Title:  "开奖日期",
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

			PushButton{
				Text:       "阳总牛逼",
				Persistent: true,
				//MinSize: Size{30, 10},
				OnClicked: func() {
					doubleabll = Randball()
					a := fmt.Sprint(doubleabll[:6])
					b := fmt.Sprint(doubleabll[6:])
					usernameTE.SetText(a)
					passwordTE.SetText(b)
					dateTE.SetText(GetDate())
					doubleabll = []int{}
				},
			},
		},
	}.Run()

}

func Randball() []int {

	rand.Seed(time.Now().Unix())

	for j := 0; j < 6; j++ {
		intn := rand.Intn(34)
		redball[intn] = intn
		doubleabll = append(doubleabll, intn)

	}
	sort.Ints(doubleabll)
	for i := 0; i < 1; i++ {
		intn := rand.Intn(17)
		redball[intn] = intn
		doubleabll = append(doubleabll, intn)
	}
	return doubleabll
}
func GetDate() string {
	now := time.Now()

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

func GetWeekday(t time.Time) string {
	Weekday := fmt.Sprint(t.Weekday())
	return Weekday
}
