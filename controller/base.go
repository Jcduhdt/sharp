package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"sharp/common/consts"
	"sharp/common/dto"
)

type BaseController struct {
}

func (bc *BaseController) ReturnJSon(c *gin.Context, data interface{}, err error) {
	resp := bc.GetRespByError(err, data)
	c.Header("sharp-header-rid", c.GetString("traceid"))
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
