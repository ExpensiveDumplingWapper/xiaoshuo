/*
 * @Descripttion: 我见青山多妩媚
 * @Date: 2022-01-04 14:32:34
 * @LastEditTime: 2022-01-10 17:37:11
 */
package index

import (
	"net/http"
	"strconv"

	"xiaoshuo/internal/services/index"
	"xiaoshuo/internal/services/raw_data"

	"github.com/gin-gonic/gin"
)

func BookSearch(ctx *gin.Context) {
	data, _ := index.BookSearch(ctx)
	menu := raw_data.GetMenus()
	page := "1"
	finalPage, _ := strconv.Atoi(page)
	nextPage := strconv.Itoa(finalPage + 1)
	keyWord := ctx.Query("searchkey")

	if ctx.GetBool("isMobile") {
		ctx.HTML(http.StatusOK, "m_book_search.tmpl", gin.H{
			"detail":           data,
			"image":            Image,
			"hostDefaultImage": HostDefaultImage,
			"menu":             menu,
			"nextPage":         nextPage,
			"keyWord":          keyWord,
		})
	} else {
		ctx.HTML(http.StatusOK, "book_search.tmpl", gin.H{
			"detail":           data,
			"image":            Image,
			"menu":             menu,
			"hostDefaultImage": HostDefaultImage,
			"nextPage":         nextPage,
			"keyWord":          keyWord,
		})
	}
}

func SearchAuthor(ctx *gin.Context) {
	data, _ := index.SearchAuthor(ctx)
	menu := raw_data.GetMenus()
	page := "1"
	finalPage, _ := strconv.Atoi(page)
	nextPage := strconv.Itoa(finalPage + 1)
	keyWord := ctx.Query("author")

	if ctx.GetBool("isMobile") {
		ctx.HTML(http.StatusOK, "m_book_search.tmpl", gin.H{
			"detail":           data,
			"image":            Image,
			"menu":             menu,
			"nextPage":         nextPage,
			"keyWord":          keyWord,
			"hostDefaultImage": HostDefaultImage,
		})
	} else {
		ctx.HTML(http.StatusOK, "book_search.tmpl", gin.H{
			"detail":           data,
			"image":            Image,
			"menu":             menu,
			"nextPage":         nextPage,
			"keyWord":          keyWord,
			"hostDefaultImage": HostDefaultImage,
		})
	}
}
