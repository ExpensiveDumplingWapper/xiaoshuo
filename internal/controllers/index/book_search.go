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
	keyWord := ctx.PostForm("searchkey")
	host := ctx.ClientIP() + ":9999"

	if ctx.GetBool("isMobile") {
		ctx.HTML(http.StatusOK, "m_book_search.tmpl", gin.H{
			"detail":     data,
			"image":      Image,
			"hostServer": host,
			"menu":       menu,
			"nextPage":   nextPage,
			"keyWord":    keyWord,
		})
	} else {
		ctx.HTML(http.StatusOK, "book_search.tmpl", gin.H{
			"detail":     data,
			"image":      Image,
			"menu":       menu,
			"hostServer": host,
			"nextPage":   nextPage,
			"keyWord":    keyWord,
		})
	}
}

func SearchAuthor(ctx *gin.Context) {
	data, _ := index.SearchAuthor(ctx)
	menu := raw_data.GetMenus()
	page := "1"
	finalPage, _ := strconv.Atoi(page)
	nextPage := strconv.Itoa(finalPage + 1)
	keyWord := ctx.PostForm("author")

	if ctx.GetBool("isMobile") {
		ctx.HTML(http.StatusOK, "m_book_search.tmpl", gin.H{
			"detail":   data,
			"image":    Image,
			"menu":     menu,
			"nextPage": nextPage,
			"keyWord":  keyWord,
		})
	} else {
		ctx.HTML(http.StatusOK, "book_search.tmpl", gin.H{
			"detail":   data,
			"image":    Image,
			"menu":     menu,
			"nextPage": nextPage,
			"keyWord":  keyWord,
		})
	}
}
