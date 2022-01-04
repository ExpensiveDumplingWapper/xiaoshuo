package middlewares

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Runtime() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		//c.Set("variable", "12345")

		// before request

		c.Next()

		// after request
		time.Since(t)
		//log.Info(latency)
	}
}
