package main

import (
	"github.com/gin-gonic/gin"
	"sharp/common/handler/log"
	"sharp/common/handler/redis"
	"sharp/controller"
	"sharp/controller/common"
	"sharp/middlewares"
)

func init(){
	log.Init()
	redis.Init()
}

func main() {
	r := gin.New()
	r.Use(gin.Recovery(),middlewares.Logger(log.Logger))

	// 用于探测服务是否正常
	r.GET("/sharp/ping", new(controller.WelcomeController).Ping)

	// common
	r.GET("/sharp/rd-test/queryCache",new(common.RdTestController).QueryCache)
	r.Run(":8000")
}
