/*
 * @Descripttion: 我见青山多妩媚
 * @Date: 2022-01-04 14:41:24
 * @LastEditTime: 2022-01-06 18:34:55
 */
package index

import (
	"xiaoshuo/internal/services/raw_data"

	"github.com/gin-gonic/gin"
)

func BookInfo(ctx *gin.Context) (res raw_data.BookInfoData, err error) {
	id := ctx.Param("id")
	res = raw_data.GetBookeInfo(id)

	return
}
