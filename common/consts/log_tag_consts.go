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

	DLTagComRequestIn    = " _com_request_in   %s"
	DLTagComRequestOut   = " _com_request_out   %s"
	DLTagCommonInfo      = " _common_info    %s"
	DLTagCommonErrorInfo = " _common_error_info    %s"

	DLTagInitRedisSuccess = " _init_redis_success    %s"
	DLTagInitRedisFailed  = " _init_redis_failed    %s"
	DLTagInitMysqlSuccess = " _init_mysql_success    %s"
	DLTagInitMysqlFailed  = " _init_mysql_failed    %s"

	DLTagGetRedisSuccess = " _get_redis_success    %s"
	DLTagGetRedisFailed  = " _get_redis_failed    %s"
	DLTagSetRedisSuccess = " _set_redis_success    %s"
	DLTagSetRedisFailed  = " _set_redis_failed    %s"
	DLTagDelRedisSuccess = " _del_redis_success    %s"
	DLTagDelRedisFailed  = " _del_redis_failed    %s"

	DLTagGetMysqlSuccess    = " _get_mysql_success    %s"
	DLTagGetMysqlFailed     = " _get_mysql_failed    %s"
	DLTagInsertMysqlSuccess = " _insert_mysql_success    %s"
	DLTagInsertMysqlFailed  = " _insert_mysql_failed    %s"
	DLTagDelMysqlSuccess    = " _del_mysql_success    %s"
	DLTagDelMysqlFailed     = " _del_mysql_failed    %s"
)
