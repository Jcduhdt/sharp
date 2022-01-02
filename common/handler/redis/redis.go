package redis

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"os"
	"sharp/common/consts"
	"sharp/common/handler/conf"
	"sharp/common/handler/log"
	"time"
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
		log.Logger.Fatalf(consts.DLTagInitRedisFailed, log.BuildLogByMap(context.Background(), logMap))
		os.Exit(1)
	}

	log.Logger.Infof(consts.DLTagInitRedisSuccess, log.BuildLogByMap(context.Background(), logMap))
	RedisClient = redisClient
}

func SetInfoWithEx(ctx context.Context, key string, expireTime int, value interface{}) (string, error) {
	logMap := map[string]interface{}{
		consts.LogCallee: "set_info_with_ex",
		"key":            key,
		"expire_time":     expireTime,
		"value":          value,
	}
	var err error
	defer func() {
		if err != nil {
			logMap[consts.LogErrMsg] = err.Error()
			log.Logger.Errorf(consts.DLTagSetRedisFailed, log.BuildLogByMap(ctx, logMap))
			return
		}
		log.Logger.Info(consts.DLTagSetRedisSuccess, log.BuildLogByMap(ctx, logMap))
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
			log.Logger.Errorf(consts.DLTagGetRedisFailed, log.BuildLogByMap(ctx, logMap))
			return
		}
		log.Logger.Info(consts.DLTagGetRedisSuccess, log.BuildLogByMap(ctx, logMap))
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
			log.Logger.Errorf(consts.DLTagDelRedisFailed, log.BuildLogByMap(ctx, logMap))
			return
		}
		log.Logger.Info(consts.DLTagDelRedisSuccess, log.BuildLogByMap(ctx, logMap))
	}()

	start := time.Now()
	result, err := redis.Int(RedisClient.Do(consts.CommandDel, key))
	cost := time.Since(start)

	logMap["cost"] = cost
	logMap["result"] = result

	return result, err
}
