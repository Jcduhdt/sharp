package rd_test

func QueryRedisCache(redisKey string) (string,error){
	return "get redis cache",nil
}

func SetRedisCache(redisKey, value string) (string,error){
	return "set redis cache",nil
}
func DelRedisCache(redisKey string) (string,error){
	return "del redis cache",nil
}
