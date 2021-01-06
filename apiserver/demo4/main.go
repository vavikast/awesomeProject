package main

import (
	"awesomeProject/apiserver/demo4/config"
	"awesomeProject/apiserver/demo4/model"
	v "awesomeProject/apiserver/demo4/pkg/version"
	"awesomeProject/apiserver/demo4/router"
	"awesomeProject/apiserver/demo4/router/middleware"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/zxmrlc/log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

var (
	cfg     = pflag.StringP("config", "c", "", "apiserver config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	pflag.Parse()
	if *version {
		v := v.Get()
		marshalIndent, err := json.MarshalIndent(&v, "", " ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(marshalIndent))
		return

	}
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
	//middlewares := []gin.HandlerFunc{}
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()
	//routes.
	router.Load(g, middleware.RequestId(), middleware.Logging())

	//strat to listening the incoming requests
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	fmt.Println(cert)
	if cert != "" && key != "" {
		go func() {
			log.Infof("Start to listening the incoming requests on https address: %s", viper.GetString("tls.addr"))
			log.Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
		}()
	}
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
