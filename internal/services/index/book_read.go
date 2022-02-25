/*
 * @Descripttion: 我见青山多妩媚
 * @Date: 2022-01-04 14:41:24
 * @LastEditTime: 2022-01-07 18:15:25
 */
package index

import (
	"strings"
	"xiaoshuo/internal/services/raw_data"

	"github.com/gin-gonic/gin"
)

func BookRead(ctx *gin.Context) (res raw_data.BookContentData, err error) {
	bookId := ctx.Param("id")
	chapterId := strings.ReplaceAll(ctx.Param("page"), "r_", "")
	res = raw_data.GetBookeRead(ctx, bookId, chapterId)

	return
}
