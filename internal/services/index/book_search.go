/*
 * @Descripttion: 我见青山多妩媚
 * @Date: 2022-01-04 14:41:24
 * @LastEditTime: 2022-01-10 17:05:44
 */
package index

import (
	"fmt"
	"xiaoshuo/internal/services/raw_data"

	"github.com/gin-gonic/gin"
)

func BookSearch(ctx *gin.Context) (res []raw_data.BookInfo, err error) {
	keyWord := ctx.PostForm("searchkey")
	// page := ctx.Query("page")
	fmt.Println(keyWord, 99999999)
	page := "1"
	res = raw_data.GetBookSearch(keyWord, page)
	return
}
