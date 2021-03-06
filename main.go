package main

import (
	"apiserver/config"
	"apiserver/model"
	"apiserver/router"
	"apiserver/router/middleware"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/ulule/limiter"
	"os"

	// _ "./docs" // docs is generated by Swag CLI, you have to import it.

	// "log"
	v "apiserver/pkg/version"
	"net/http"
	"time"

	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	mgin "github.com/ulule/limiter/drivers/middleware/gin"
	sredis "github.com/ulule/limiter/drivers/store/redis"
)

var (
	cfg     = pflag.StringP("config", "c", "", "apiserver config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	//命令行读取
	pflag.Parse()
	if *version {
		v := v.Get()
		marshaled, err := json.MarshalIndent(&v, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshaled))
		return
	}
	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	// init DB
	model.DB.Init()
	defer model.DB.Close()

	//init limiter
	limiterM,_:=initLimiter()
	// 根据配置文件设置gin的运行模式
	gin.SetMode(viper.GetString("run_mode"))

	g := gin.New()

	router.Load(
		g,
		middleware.RequestID(),
		middleware.Logging(),
		limiterM,
	)
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	// tls
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			log.Infof("Start to listening the incoming requests on https address: %s", viper.GetString("tls.addr"))
			log.Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
		}()
	}
	// http
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
	return errors.New("Cannot connect to the router.")
}

func initLimiter() (limiterM gin.HandlerFunc, err error) {
	// Define a limit rate to 4 requests per hour.
	rate, err := limiter.NewRateFromFormatted(viper.GetString("limiter.global"))
	if err != nil {
		log.Fatal("limiter.NewRateFromFormatted err:%v", err)
		return
	}

	// Create a redis client.
	option, err := redis.ParseURL("redis://localhost:6379/0")
	if err != nil {
		log.Fatal("limiter.NewRateFromFormatted err:%v", err)
		return
	}
	client := redis.NewClient(option)
	log.Info("redis is connected")
	// Create a store with the redis client.
	store, err := sredis.NewStoreWithOptions(client, limiter.StoreOptions{
		Prefix:   "limiter_gin_example",
		MaxRetry: 3,
	})
	if err != nil {
		log.Fatal("sredis.NewStoreWithOptions err:%v", err)
		return
	}

	// Create a new middleware with the limiter instance.
	limiterM = mgin.NewMiddleware(limiter.New(store, rate))
	return
}
