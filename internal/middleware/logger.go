package middleware

import (
	"EnGin/internal/global"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		traceID := c.GetString("trace_id")

		if raw != "" {
			path = path + "?" + raw
		}

		// 状态码颜色
		statusColor := getStatusColor(statusCode)
		methodColor := getMethodColor(method)
		resetColor := "\033[0m"

		logMsg := fmt.Sprintf("%s %3d %s| %13v | %10s | %s%-5s%s %s | trace_id:%s",
			statusColor, statusCode, resetColor,
			latency,
			clientIP,
			methodColor, method, resetColor,
			path,
			traceID,
		)

		if statusCode >= 500 {
			global.Log.Error(logMsg)
		} else {
			global.Log.Info(logMsg)
		}
	}
}

func getStatusColor(code int) string {
	switch {
	case code >= 200 && code < 300:
		return "\033[97;42m" // 绿底白字
	case code >= 300 && code < 400:
		return "\033[90;47m" // 白底黑字
	case code >= 400 && code < 500:
		return "\033[97;43m" // 黄底白字
	default:
		return "\033[97;41m" // 红底白字
	}
}

func getMethodColor(method string) string {
	switch method {
	case "GET":
		return "\033[97;44m" // 蓝底白字
	case "POST":
		return "\033[97;46m" // 青底白字
	case "PUT":
		return "\033[97;43m" // 黄底白字
	case "DELETE":
		return "\033[97;41m" // 红底白字
	case "PATCH":
		return "\033[97;42m" // 绿底白字
	case "HEAD":
		return "\033[97;45m" // 紫底白字
	case "OPTIONS":
		return "\033[90;47m" // 白底黑字
	default:
		return "\033[0m"
	}
}
