package info

import (
	"time"

	"github.com/JiHanHuang/gin_vue/pkg/logging"
	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func MSG() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		status := c.Writer.Status()
		logging.Info(
			"path:", path,
			"method:", method,
			"status:", status,
			"client_ip:", clientIP,
			"latency:", latency,
		)
	}
}
