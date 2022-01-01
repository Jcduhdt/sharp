package common

import (
	"github.com/gin-gonic/gin"
	"sharp/common/consts"
	"sharp/common/dto"
	"sharp/common/handler/log"
	"sharp/controller"
	"sharp/model/rd_test"
)

type RdTestController struct {
	controller.BaseController
}

func (rtc *RdTestController) QueryCache(ctx *gin.Context) {
	logMap := map[string]interface{}{
		consts.LogCallee: "query_cache",
	}
	var err error
	var data interface{}
	defer func() {
		if err != nil {
			data = err.Error()
			logMap[consts.LogErrMsg] = err.Error()
			log.Logger.Errorf("s", log.BuildLogByMap(ctx, logMap))
		}
		rtc.ReturnJSon(ctx, data, err)
	}()


	var redisReq dto.RedisReq
	err = ctx.ShouldBindUri(&redisReq)
	if err != nil {
		return
	}
	res, err := rd_test.QueryRedisCache(redisReq.RedisKey)
	if err != nil {
		return
	}
	data = res
}
