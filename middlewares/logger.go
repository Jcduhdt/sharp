package middlewares

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"sharp/common/consts"
	"sharp/common/dto"
	"sharp/common/handler/log"
	"strings"
	"time"
)

type BodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w BodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w BodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// Logger 接收gin框架默认的日志
func Logger(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// trace
		traceId := strings.Replace(uuid.NewV4().String(), "-", "", -1)
		c.Set("traceid", traceId)

		// 请求体
		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		requestBody, _ := ioutil.ReadAll(tee)
		c.Request.Body = ioutil.NopCloser(&buf)

		bodyLogWriter := &BodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		// 使用下一个中间件
		c.Next()

		// 载入响应内容
		responseBody := bodyLogWriter.body.Bytes()
		response := dto.Response{}
		if len(responseBody) > 0 {
			_ = json.Unmarshal(responseBody, &response)
		}

		logMap := map[string]interface{}{
			"host":         c.ClientIP(),
			"user-agent":   c.Request.UserAgent(),
			"method":       c.Request.Method,
			"path":         c.Request.URL.Path,
			"content_type": c.Request.Header.Get("Content-Type"),
			"proto":        c.Request.Proto,
			"req_params":   c.Request.URL.RawQuery,
			"req_body":     string(requestBody),
			"resp":         response,
			"cost":         time.Since(start),
		}

		logger.Infof(consts.DLTagComRequestOut, log.BuildLogByMap(c, logMap))
	}
}
