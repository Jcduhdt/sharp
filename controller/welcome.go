package controller

import (
	"github.com/gin-gonic/gin"
)

type WelcomeController struct {
	BaseController
}

func (wc *WelcomeController) Ping(c *gin.Context) {
	data := "pong"
	wc.ReturnJSon(c, data, nil)
}
