package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"sharp/common/consts"
	"sharp/common/dto"
)

type BaseController struct {
}

func (bc *BaseController) ReturnJSon(c *gin.Context, data interface{}, err error) {
	resp := bc.GetRespByError(err, data)
	resp.TraceId = c.GetString("traceid")
	c.JSON(http.StatusOK, resp)
}

func (bc *BaseController) GetRespByError(err error, data interface{}) dto.Response {
	errMsg := ""
	errNo := 0

	if err != nil {
		if errors.Is(err, consts.ErrParamsInvalid) {
			errNo = consts.RespParamsInvalidCode
		} else {
			errNo = consts.RespFailedCode
		}
	}

	errMsg = consts.RespMsg[errNo]

	return dto.Response{
		Errno:  errNo,
		Errmsg: errMsg,
		Data:   data,
	}
}
