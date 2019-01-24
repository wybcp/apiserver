package main

import (
	"apiserver/config"
	"apiserver/model"
	"apiserver/router"
	"errors"

	// "log"
	"net/http"
	"time"

	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
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
	// init DB
	model.DB.Init()
	defer model.DB.Close()

	// 根据配置文件设置gin的运行模式
	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()
	middlewares := []gin.HandlerFunc{}
	router.Load(
		g,
		middlewares...,
	)
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()
	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router")
}
