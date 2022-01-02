package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"sharp/common/consts"
	"sharp/common/handler/log"
	"strings"
	"time"
)

// Logger 接收gin框架默认的日志
func Logger(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		traceId := strings.Replace(uuid.NewV4().String(), "-", "", -1)
		c.Set("traceid", traceId)

		c.Next()

		logMap := map[string]interface{}{
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"req_params": c.Request.URL.RawQuery,
			"req_body":   c.Request.PostForm.Encode(),
			"host":       c.ClientIP(),
			"user-agent": c.Request.UserAgent(),
			"cost":       time.Since(start),
		}

		logger.Infof(consts.DLTagComRequestOut, log.BuildLogByMap(c, logMap))
	}
}
