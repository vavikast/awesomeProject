package main

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type Config struct {
	AppName  string      `ini:"app_name"`
	LogLevel string      `ini:"log_level"`
	MySQL    MySQLConfig `ini:"mysql"`
	Redis    RedisConfig `ini:"redis"`
}

type MySQLConfig struct {
	IP       string `ini:"ip"`
	Port     string `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type RedisConfig struct {
	IP   string `ini:"ip"`
	Port int    `ini:"port"`
}

func main() {
	cfg, err := ini.Load("./ES/httpweb/server/my.ini")
	if err != nil {
		fmt.Println("load my.ini failed:", err)
	}
	c := Config{}
	cfg.MapTo(&c)
	fmt.Println(c)

}
