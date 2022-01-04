package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Debug() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//fmt.Println("RemoteIP", request.GetRemoteIP(ctx))
		//fmt.Println("ClientIP", ctx.ClientIP())
		//fmt.Println("RemoteAddr", ctx.Request.RemoteAddr)
		//fmt.Println("===============")
		//for hk, hv := range ctx.Request.Header {
		//	fmt.Println(hk, hv)
		//}
		//fmt.Println("===============")
	}
}
