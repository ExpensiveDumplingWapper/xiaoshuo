package middlewares

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"runtime"
	"strings"

	"xiaoshuo/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

const RecoveryLogNamed = "recovery"

// RecoveryWithZap returns a gin.HandlerFunc (middleware)
// that recovers from any panics and logs requests using uber-go/zap.
// All errors are logged using zap.Error().
// stack means whether output the stack info.
// The stack info is easy to find where the error occurs but the stack info is too large.
func RecoveryWithZap() gin.HandlerFunc {
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
				if brokenPipe {
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}
				buf := make([]byte, 2048)
				n := runtime.Stack(buf, false)
				stackInfo := fmt.Sprintf("%s", buf[:n])
				fmt.Println(stackInfo)
				if message, ok := err.(string); ok {
					logger.Error("panic_error", "ServerPanicMessage: "+message+"\nStackInfo:\n"+stackInfo, c)
					fmt.Println(message)
				} else if message, ok := err.(runtime.Error); ok {
					logger.Error("panic_error", "ServerPanicMessage: "+message.Error()+"\nStackInfo:\n"+stackInfo, c)
					fmt.Println(message.Error())
				} else {
					logger.Error("panic_error", "ServerPanicMessage:"+" StackInfo:"+stackInfo, c)
				}
				c.JSON(http.StatusInternalServerError, gin.H{"code": 520, "message": "Internal Server Error"})
			}
		}()
		c.Next()
	}
}
