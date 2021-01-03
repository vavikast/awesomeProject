package main

import (
	"awesomeProject/apiserver/demo4/config"
	"awesomeProject/apiserver/demo4/model"
	"awesomeProject/apiserver/demo4/router"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/zxmrlc/log"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()
	//init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	//init db
	model.DB.Init()
	defer model.DB.Close()

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	//create The Gin engine
	g := gin.New()

	//gin middlewars
	middlewares := []gin.HandlerFunc{}
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()
	//routes.
	router.Load(g, middlewares...)
	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

func pingServer() error {
	for i := 0; i < 2; i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
	}
	// Sleep for a second to continue the next ping.
	log.Info("Waiting for the router, retry in 1 second.")
	time.Sleep(time.Second)
	return errors.New("Cannot connect to the router")
}
