/*
 * @Descripttion: 我见青山多妩媚
 * @Date: 2022-01-10 17:46:44
 * @LastEditTime: 2022-01-10 18:03:30
 */
package middlewares

import (
	"fmt"
	"xiaoshuo/pkg/request"

	"github.com/gin-gonic/gin"
)

func ParseUa() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ua := request.HttpParseUa(ctx)
		ctx.Set("ua", ua)

		fmt.Println(ua.Platform)
		fmt.Println("<br><------------------------><br>")
		fmt.Println(ctx.GetHeader("User-Agent"))
		ctx.Next()
	}
}
