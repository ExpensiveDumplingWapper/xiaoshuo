/*
 * @Descripttion: 我见青山多妩媚
 * @Date: 2022-01-04 14:41:24
 * @LastEditTime: 2022-01-10 17:36:57
 */
package index

import (
	"xiaoshuo/internal/services/raw_data"

	"github.com/gin-gonic/gin"
)

func BookSearch(ctx *gin.Context) (res []raw_data.BookInfo, err error) {
	keyWord := ctx.PostForm("searchkey")
	// page := ctx.Query("page")
	// fmt.Println(keyWord, )
	page := "1"
	res = raw_data.GetBookSearch(keyWord, page)
	return
}
