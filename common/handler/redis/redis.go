package redis

import (
	"context"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"

	"sharp/common/consts"
	"sharp/common/handler/conf"
	"sharp/common/handler/log"
)

var (
	RedisClient redis.Conn
)

func Init() {

	addrs := conf.Viper.GetString("redis.addrs")
	auth := conf.Viper.GetString("redis.auth")
	readTimeOut := time.Duration(conf.Viper.GetInt("redis.read_timeout")) * time.Millisecond
	writeTimeOut := time.Duration(conf.Viper.GetInt("redis.write_timeout")) * time.Millisecond
	connTimeOut := time.Duration(conf.Viper.GetInt("redis.conn_timeout")) * time.Millisecond

	setPwd := redis.DialPassword(auth)
	setRto := redis.DialReadTimeout(readTimeOut)
	setWto := redis.DialReadTimeout(writeTimeOut)
	setCto := redis.DialConnectTimeout(connTimeOut)

	redisClient, err := redis.Dial("tcp", addrs, setPwd, setRto, setWto, setCto)

	logMap := map[string]interface{}{
		"addrs": addrs,
	}
	if err != nil {
		logMap[consts.LogErrMsg] = err.Error()
		log.FatalMap(context.Background(), consts.DLTagInitRedisFailed, logMap)
		os.Exit(1)
	}
	log.InfoMap(context.Background(), consts.DLTagInitRedisSuccess, logMap)
	RedisClient = redisClient
}

func SetInfoWithEx(ctx context.Context, key string, expireTime int, value interface{}) (string, error) {
	logMap := map[string]interface{}{
		consts.LogCallee: "set_info_with_ex",
		"key":            key,
		"expire_time":    expireTime,
		"value":          value,
	}
	var err error
	defer func() {
		if err != nil {
			logMap[consts.LogErrMsg] = err.Error()
			log.ErrorMap(ctx, consts.DLTagSetRedisFailed, logMap)
			return
		}
		log.InfoMap(ctx, consts.DLTagSetRedisSuccess, logMap)
	}()

	if expireTime <= 0 {
		expireTime = consts.DefaultExpireTime
	}

	start := time.Now()
	result, err := redis.String(RedisClient.Do(consts.CommandSetEx, key, expireTime, value))
	cost := time.Since(start)

	logMap["cost"] = cost
	logMap["result"] = result

	return result, err
}

func GetInfo(ctx context.Context, key string) (string, error) {
	logMap := map[string]interface{}{
		consts.LogCallee: "set_info_with_ex",
		"key":            key,
	}
	var err error
	defer func() {
		if err != nil {
			logMap[consts.LogErrMsg] = err.Error()
			log.ErrorMap(ctx, consts.DLTagGetRedisFailed, logMap)
			return
		}
		log.InfoMap(ctx, consts.DLTagGetRedisSuccess, logMap)
	}()

	start := time.Now()
	result, err := redis.String(RedisClient.Do(consts.CommandGet, key))
	cost := time.Since(start)

	logMap["cost"] = cost
	logMap["result"] = result

	return result, err
}

func DelInfo(ctx context.Context, key string) (int, error) {
	logMap := map[string]interface{}{
		consts.LogCallee: "set_info_with_ex",
		"key":            key,
	}
	var err error
	defer func() {
		if err != nil {
			logMap[consts.LogErrMsg] = err.Error()
			log.ErrorMap(ctx, consts.DLTagDelRedisFailed, logMap)
			return
		}
		log.InfoMap(ctx, consts.DLTagDelRedisSuccess, logMap)
	}()

	start := time.Now()
	result, err := redis.Int(RedisClient.Do(consts.CommandDel, key))
	cost := time.Since(start)

	logMap["cost"] = cost
	logMap["result"] = result

	return result, err
}
