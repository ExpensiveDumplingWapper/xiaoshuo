/*
 * @Descripttion: 我见青山多妩媚
 * @Date: 2022-01-04 14:41:24
 * @LastEditTime: 2022-01-06 19:04:26
 */
package index

import (
	"xiaoshuo/internal/services/raw_data"

	"github.com/gin-gonic/gin"
)

func BookRead(ctx *gin.Context) (res raw_data.BookInfoData, err error) {
	bookId := ctx.Param("bookid")
	chapterId := ctx.Param("chapterid")
	res = raw_data.GetBookeRead(bookId, chapterId)

	return
}
