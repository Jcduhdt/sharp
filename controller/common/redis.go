package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	redigo "github.com/gomodule/redigo/redis"
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
		if err != nil{
			data = err.Error()
			logMap[consts.LogErrMsg] = err.Error()
			log.Logger.Errorf(consts.DLTagCommonErrorInfo, log.BuildLogByMap(ctx, logMap))
		}
		rtc.ReturnJSon(ctx, data, err)
	}()

	var req dto.RedisReq
	err = ctx.BindQuery(&req)
	if err != nil {
		return
	}

	res, err := rd_test.QueryRedisCache(ctx, req.RedisKey)

	if errors.Is(err,redigo.ErrNil){
		data = err.Error()
		err = nil
		return
	}

	data = res
}


func (rtc *RdTestController) DelCache(ctx *gin.Context) {
	logMap := map[string]interface{}{
		consts.LogCallee: "del_cache",
	}
	var err error
	var data interface{}
	defer func() {
		if err != nil {
			data = err.Error()
			logMap[consts.LogErrMsg] = err.Error()
			log.Logger.Errorf(consts.DLTagCommonErrorInfo, log.BuildLogByMap(ctx, logMap))
		}
		rtc.ReturnJSon(ctx, data, err)
	}()

	var req dto.RedisReq
	err = ctx.BindQuery(&req)
	if err != nil {
		return
	}
	res, err := rd_test.DelRedisCache(ctx, req.RedisKey)
	if err != nil {
		return
	}
	data = res
}


func (rtc *RdTestController) SetCacheWithEx(ctx *gin.Context) {
	logMap := map[string]interface{}{
		consts.LogCallee: "set_cache_with_ex",
	}
	var err error
	var data interface{}
	defer func() {
		if err != nil {
			data = err.Error()
			logMap[consts.LogErrMsg] = err.Error()
			log.Logger.Errorf(consts.DLTagCommonErrorInfo, log.BuildLogByMap(ctx, logMap))
		}
		rtc.ReturnJSon(ctx, data, err)
	}()

	var req dto.RedisSetReq
	err = ctx.BindJSON(&req)
	if err != nil {
		return
	}
	res, err := rd_test.SetRedisCache(ctx, req.RedisKey, req.ExpireTime, req.Value)
	if err != nil {
		return
	}
	data = res
}
