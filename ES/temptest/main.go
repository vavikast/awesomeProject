package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var (
	redball    = make(map[int]int)
	bullball   = make(map[int]int)
	doubleabll = make([]int, 0)
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(Randball())
	}

}
func Randball() []int {
	rand.Seed(time.Now().Unix())
	doubleabll = []int{}
	//redball = map[int]int{}
	for j := 0; j < 6; j++ {
		intn := rand.Intn(33)
		fmt.Println(intn)
		fmt.Println(j)
		if _, ok := redball[intn]; ok {
			j--
			fmt.Println(redball)
		} else {
			redball[intn] = intn + 1
			doubleabll = append(doubleabll, redball[intn])
			fmt.Println(doubleabll)
		}

	}
	sort.Ints(doubleabll)
	for i := 0; i < 1; i++ {
		intn := rand.Intn(16)
		if _, ok := bullball[intn]; ok {
			i--
		} else {
			bullball[intn] = intn + 1
			doubleabll = append(doubleabll, bullball[intn])
			fmt.Println(doubleabll)
		}

	}
	return doubleabll
}
