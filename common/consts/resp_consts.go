package consts

const (
	// 响应成功
	RespSuccessCode = 0
	// 参数错误
	RespParamsInvalidCode = 1
	// 服务内部错误
	RespFailedCode = 2
)

var RespMsg = map[int]string{
	RespSuccessCode:       "Success",
	RespParamsInvalidCode: "Request Params Invalid",
	RespFailedCode:        "NetWork Error",
}
