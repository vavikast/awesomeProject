package main

import (
	"bytes"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config")             // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")               // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("./goredis/goviper/") // 查找配置文件所在的路径
	err := viper.ReadInConfig()               // 查找并读取配置文件
	if err != nil {                           // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))

	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)

	})
	word := viper.Get("name")
	fmt.Println(word)
	viper.ReadConfig(bytes.NewBuffer([]byte("nebwbbw")))
	b := byte()
}
