package main

import (
	"awesomeProject/goredis/common"
	"fmt"
	"time"
)

type Inventory struct {
	Id     string
	Data   string
	Cached int64
}

func NewInventory(id, data string, cached int64) Inventory {
	return Inventory{
		Id:     id,
		Data:   data,
		Cached: cached,
	}
}

func Get(id string) Inventory {
	return NewInventory(id, "data to cache...", time.Now().Unix())
}
func main() {
	rdb := common.InitClient()
	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("It's ok")
}

func Check_token(conn, token string) (user string) {
	return user

}
