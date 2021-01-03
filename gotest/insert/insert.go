package main

import (
	"awesomeProject/gotest/insert/config"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	for {
		fmt.Println(viper.GetString("runmode"))
		fmt.Println(viper.GetString("addr"))
		fmt.Println(viper.GetString("name"))
		fmt.Println(viper.GetString("url"))
		time.Sleep(4 * time.Second)
	}

}
