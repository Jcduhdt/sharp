package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"

	"sharp/common/handler/conf"
	"sharp/common/handler/env"
	"sharp/common/handler/log"
	"sharp/common/handler/mysql"
	"sharp/common/handler/redis"
	"sharp/controller"
	"sharp/controller/common"
	"sharp/middlewares"
)

var (
	confPath string
)

func init() {
	flag.StringVar(&confPath, "c", "./conf/", "-c set config file path")
	flag.Parse()
	fmt.Printf("confPath is %s\n", confPath)

	env.Init()
	conf.InitConf(confPath)
	log.Init()
	redis.Init()
	mysql.Init()
}

func main() {
	r := gin.New()
	r.Use(gin.Recovery(), middlewares.Logger(log.Logger))

	// 用于探测服务是否正常
	r.GET("/sharp/ping", new(controller.WelcomeController).Ping)

	// common
	r.GET("/sharp/rd-test/queryCache", new(common.RdTestController).QueryCache)
	r.GET("/sharp/rd-test/deleteCache", new(common.RdTestController).DelCache)
	r.POST("/sharp/rd-test/setCacheWithEx", new(common.RdTestController).SetCacheWithEx)

	r.POST("/sharp/login", new(common.LoginController).Login)
	r.POST("/sharp/register", new(common.LoginController).Register)

	r.Run(conf.Viper.GetString("server.port"))
}
