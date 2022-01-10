/*
 * @Descripttion: 我见青山多妩媚
 * @Date: 2022-01-04 14:41:24
 * @LastEditTime: 2022-01-10 14:51:48
 */
package index

import (
	"strconv"
	"xiaoshuo/internal/services/raw_data"

	"github.com/gin-gonic/gin"
)

func BookType(ctx *gin.Context) (res, hotRes []raw_data.BookInfo, err error) {
	bookType := ctx.Param("booktype")
	page := ctx.Param("paga")
	finalPage, _ := strconv.Atoi(page)
	if finalPage <= 1 {
		page = "1"
	}
	res = raw_data.GetBookeType(bookType, page)
	hotRes = raw_data.GetBookeTypeHot(bookType)
	return
}
