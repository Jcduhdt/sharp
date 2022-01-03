package handler

import (
	"sharp/common/handler/conf"
	"sharp/common/handler/env"
	"sharp/common/handler/log"
	"sharp/common/handler/mysql"
	"sharp/common/handler/redis"
	"sync"
)

var once sync.Once

func EnvInitTest() {
	once.Do(func() {
		env.Init()
		conf.InitConf(conf.FindConfigDir())
		log.Init()
		mysql.Init()
		redis.Init()
	})
}
