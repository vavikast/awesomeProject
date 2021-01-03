package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)
import "awesomeProject/apiserver/demo1/router"

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()

	//init config

	//create The Gin engine
	g := gin.New()

	//gin middlewars
	middlewares := []gin.HandlerFunc{}
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()
	//routes.
	router.Load(g, middlewares...)
	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	log.Println(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

func pingServer() error {
	for i := 0; i < 2; i++ {
		resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
	}
	// Sleep for a second to continue the next ping.
	log.Print("Waiting for the router, retry in 1 second.")
	time.Sleep(time.Second)
	return errors.New("Cannot connect to the router")
}
