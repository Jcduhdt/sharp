package rd_test

import (
	"context"
	"encoding/json"

	"sharp/common/handler/redis"
)

func QueryRedisCache(ctx context.Context, redisKey string) (string, error) {
	return redis.GetInfo(ctx, redisKey)
}

func SetRedisCache(ctx context.Context, redisKey string, expireTime int, value interface{}) (string, error) {
	valBytes, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return redis.SetInfoWithEx(ctx, redisKey, expireTime, valBytes)
}
func DelRedisCache(ctx context.Context, redisKey string) (int, error) {
	return redis.DelInfo(ctx, redisKey)
}
