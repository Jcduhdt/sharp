package common

import (
	"github.com/gin-gonic/gin"
	"sharp/common/consts"
	"sharp/common/dto"
	"sharp/common/handler/log"
	"sharp/controller"
	"sharp/model/user_service"
)

type LoginController struct {
	controller.BaseController
}

func (lc *LoginController) Login(ctx *gin.Context) {
	logMap := map[string]interface{}{
		consts.LogCallee: "login_controller",
	}
	var err error
	var data interface{}
	defer func() {
		if err != nil {
			logMap[consts.LogErrMsg] = err.Error()
			log.ErrorMap(ctx, consts.DLTagCommonErrorInfo, logMap)
		}
		lc.ReturnJSon(ctx, data, err)
	}()

	var req dto.LoginAndRegisterReq
	err = ctx.BindJSON(&req)
	if err != nil {
		return
	}

	loginStatus, err := user_service.Login(ctx, req)

	data = loginStatus
}

func (lc *LoginController) Register(ctx *gin.Context) {
	logMap := map[string]interface{}{
		consts.LogCallee: "register_controller",
	}
	var err error
	var data interface{}
	defer func() {
		if err != nil {
			logMap[consts.LogErrMsg] = err.Error()
			log.ErrorMap(ctx, consts.DLTagCommonErrorInfo, logMap)
		}
		lc.ReturnJSon(ctx, data, err)
	}()

	var req dto.LoginAndRegisterReq
	err = ctx.BindJSON(&req)
	if err != nil {
		return
	}

	loginStatus, err := user_service.Register(ctx, req)

	data = loginStatus
}
