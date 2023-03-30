package middware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// LoggerMiddleware 接收gin框架默认的日志
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		zapLogFields := make([]zap.Field, 0)
		startTime := time.Now()
		c.Next() // 调用该请求的剩余处理程序
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000))))
		statusCode := c.Writer.Status()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		url := c.Request.RequestURI
		zapLogFields = append(zapLogFields, zap.String("SpendTime", spendTime))
		zapLogFields = append(zapLogFields, zap.String("path", url))
		zapLogFields = append(zapLogFields, zap.String("Method", method))
		zapLogFields = append(zapLogFields, zap.Int("status", statusCode))

		opts := zap.Fields(zapLogFields...)
		log := zap.L()
		log.WithOptions(opts)
		if len(c.Errors) > 0 { // 创建内部错误
			c.Errors.ByType(gin.ErrorTypePrivate)
		}
		if statusCode >= 500 {
			log.Error("[Error]")
		} else if statusCode >= 400 {
			log.Warn("[Warning]")
		} else {
			log.Info("[Info]")
		}
	}
}

func RecoveryLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				log := zap.L()
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					log.Error(c.Request.URL.Path,
						zap.Any("[Error]", err),
						zap.String("[Request]", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				log.Error("[Recovery from panic]",
					zap.Any("[Error]", err),
					zap.String("[Request}", string(httpRequest)),
					zap.String("[Stack]", string(debug.Stack())),
				)

				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
