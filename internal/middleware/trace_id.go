package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Set("trace_id", requestID)
		c.Header("X-Trace-ID", requestID)
		c.Next()
	}
}
