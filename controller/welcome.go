package controller

import (
	"github.com/gin-gonic/gin"
	"sharp/common/handler/log"
)

type WelcomeController struct {
	BaseController
}

func (wc *WelcomeController) Ping(c *gin.Context) {
	log.Logger.Infof("%s",log.BuildLogByMap(c,map[string]interface{}{"test":"zx"}))
	data := "pong"
	wc.ReturnJSon(c, data, nil)
}
