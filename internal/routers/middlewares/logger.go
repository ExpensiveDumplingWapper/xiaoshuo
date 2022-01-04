package middlewares

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"xiaoshuo/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

func Log() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()

		statusCode := ctx.Writer.Status()
		spend := fmt.Sprintf(" spend: %fs", time.Since(startTime).Seconds())

		switch {
		case statusCode >= 400 && statusCode <= 499:
			logger.Warn("warn", "statusCode: "+strconv.Itoa(statusCode)+spend, ctx)
		case statusCode >= 500:
			logger.Error("error", "statusCode: "+strconv.Itoa(statusCode)+spend, ctx)
		default:
			returnAids, exists := ctx.Get("returnAids")
			if exists == false {
				logger.Empty(ctx)
			} else {
				s, _ := json.Marshal(returnAids)
				logger.Success(string(s)+spend, ctx)
			}
		}
	}
}
