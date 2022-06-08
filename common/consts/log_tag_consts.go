package consts

const (
	// 调用函数
	LogCallee = "callee"
	// 日志分割符
	LogParamDelimiter = "||"
	// key=val
	LogEqualSymbol = "="
	// 错误信息
	LogErrMsg = "err_msg"
	// 调用入参
	LogParams = "params"
	// 请求入参
	LogReqParams = "req_params"

	LogTemplateSuffix = "    %s"

	DLTagComRequestIn    = " _com_request_in"
	DLTagComRequestOut   = " _com_request_out"
	DLTagCommonInfo      = " _common_info"
	DLTagCommonErrorInfo = " _common_error_info"

	DLTagInitRedisSuccess = " _init_redis_success"
	DLTagInitRedisFailed  = " _init_redis_failed"
	DLTagInitMysqlSuccess = " _init_mysql_success"
	DLTagInitMysqlFailed  = " _init_mysql_failed"

	DLTagGetRedisSuccess = " _get_redis_success"
	DLTagGetRedisFailed  = " _get_redis_failed"
	DLTagSetRedisSuccess = " _set_redis_success"
	DLTagSetRedisFailed  = " _set_redis_failed"
	DLTagDelRedisSuccess = " _del_redis_success"
	DLTagDelRedisFailed  = " _del_redis_failed"

	DLTagGetMysqlSuccess    = " _get_mysql_success"
	DLTagGetMysqlFailed     = " _get_mysql_failed"
	DLTagInsertMysqlSuccess = " _insert_mysql_success"
	DLTagInsertMysqlFailed  = " _insert_mysql_failed"
	DLTagDelMysqlSuccess    = " _del_mysql_success"
	DLTagDelMysqlFailed     = " _del_mysql_failed"
)
