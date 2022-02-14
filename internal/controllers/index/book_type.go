/*
 * @Descripttion: 我见青山多妩媚
 * @Date: 2022-01-04 14:32:34
 * @LastEditTime: 2022-01-10 14:52:12
 */
package index

import (
	"net/http"
	"strconv"

	"xiaoshuo/internal/services/index"
	"xiaoshuo/internal/services/raw_data"

	"github.com/gin-gonic/gin"
)

func BookType(ctx *gin.Context) {

	data, hotData, _ := index.BookType(ctx)
	menu := raw_data.GetMenus()

	bookType := ctx.Param("booktype")
	page := ctx.Param("paga")
	finalPage, _ := strconv.Atoi(page)

	nextPage := strconv.Itoa(finalPage + 1)

	prevPageTmp := finalPage - 1
	prevPage := "1"
	if prevPageTmp > 1 {
		prevPage = strconv.Itoa(prevPageTmp)
	}
	if ctx.GetBool("isMobile") {
		ctx.HTML(http.StatusOK, "m_book_type.tmpl", gin.H{
			"detail":           data,
			"hotData":          hotData,
			"image":            Image,
			"hostDefaultImage": HostDefaultImage,
			"menu":             menu,
			"bookType":         bookType,
			"nextPage":         nextPage,
			"prevPage":         prevPage,
		})
	} else {
		ctx.HTML(http.StatusOK, "book_type.tmpl", gin.H{
			"detail":           data,
			"hotData":          hotData,
			"image":            Image,
			"hostDefaultImage": HostDefaultImage,
			"menu":             menu,
			"bookType":         bookType,
			"nextPage":         nextPage,
			"prevPage":         prevPage,
		})
	}

}
