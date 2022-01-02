package dto

// common redis
type RedisReq struct {
	RedisKey string `json:"redis_key" form:"redis_key" binding:"required"`
}

type RedisSetReq struct {
	RedisKey   string      `json:"redis_key" form:"redis_key" binding:"required"`
	ExpireTime int         `json:"expire_time"`
	Value      interface{} `json:"value" form:"value" binding:"required"`
}
