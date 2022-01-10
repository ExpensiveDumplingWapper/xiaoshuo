/*
 * @Descripttion: 我见青山多妩媚
 * @Date: 2022-01-10 17:46:44
 * @LastEditTime: 2022-01-10 19:05:15
 */
package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func ParseUa() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ua := request.HttpParseUa(ctx)
		// ctx.Set("ua", ua)

		// fmt.Println(ua.Platform)
		// fmt.Println("<br><------------------------><br>")
		// fmt.Println()

		isMobile := false
		UA := ctx.GetHeader("User-Agent")

		if find := strings.Contains(UA, "Mobile"); find {
			isMobile = true
		}
		if find := strings.Contains(UA, "UCBrowser"); find {
			isMobile = true
		}

		ctx.Set("isMobile", isMobile)
		ctx.Next()
	}
}
