package main

import (
	"awesomeProject/ES/blog-service/global"
	"awesomeProject/ES/blog-service/internal/model"
	"awesomeProject/ES/blog-service/internal/routers"
	"awesomeProject/ES/blog-service/pkg/logger"
	"awesomeProject/ES/blog-service/pkg/setting"
	"github.com/natefinch/lumberjack"
	"log"
	"net/http"
	"time"
)

func init()  {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v",err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v",err)
	}

}

func setupLogger() error{
	logger.NewLogger(&lumberjack.Logger{
		Filename:   global.AppSetting.LogSavePath+"/"+global.AppSetting.LogFileName+global.AppSetting.LogFileExt,
		MaxSize:    600,
		MaxAge:     10,
		LocalTime:  true,
	},"",log.LstdFlags).WithCaller(2)
	return nil
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return  err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return  err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT",&global.JWTSetting)
	if err != nil{
		return err
	}
	err = setting.ReadSection("Email",&global.EmailSetting)
	if err != nil{
		return err
	}
	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil

}
// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/go-programming-tour-book
func main()  {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:              ":8080",
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       10*time.Second,
		WriteTimeout:      10*time.Second,
		MaxHeaderBytes:    1<<20,
	}

	s.ListenAndServe()

}

func setupDBEngine() error  {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}
