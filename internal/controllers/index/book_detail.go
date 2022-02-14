/*
 * @Descripttion: 我见青山多妩媚
 * @Date: 2022-01-04 14:32:34
 * @LastEditTime: 2022-01-11 18:32:44
 */
package index

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"xiaoshuo/internal/services/index"
	"xiaoshuo/internal/services/raw_data"

	"github.com/gin-gonic/gin"
)

type ChapterList struct {
	Page string
	Text string
}

func BookInfo(ctx *gin.Context) {

	pageType := ctx.Param("page")
	menu := raw_data.GetMenus()
	// page 页含有c 就是于都章节页面    含有r 是阅读页面
	if strings.HasPrefix(pageType, "c_") {
		data, _ := index.BookInfo(ctx)
		chapterList := make([]ChapterList, 0)
		for i := 0; i < data.TotalPage; i++ {
			chapterList = append(chapterList, ChapterList{Page: strconv.Itoa(i + 1), Text: fmt.Sprintf("第%s-第%s条", strconv.Itoa(1+i*30), strconv.Itoa(30+i*30))})
		}
		page := strings.ReplaceAll(pageType, "c_", "")
		finalPage, _ := strconv.Atoi(page)
		nextPage := strconv.Itoa(finalPage + 1)
		prevPageTmp := finalPage - 1
		prevPage := "1"
		if prevPageTmp > 1 {
			prevPage = strconv.Itoa(prevPageTmp)
		}
		if ctx.GetBool("isMobile") {
			ctx.HTML(http.StatusOK, "m_book_detail.tmpl", gin.H{
				"detail":           data,
				"image":            Image,
				"menu":             menu,
				"chapterList":      chapterList,
				"nextPage":         nextPage,
				"prevPage":         prevPage,
				"hostDefaultImage": HostDefaultImage,
			})
		} else {
			ctx.HTML(http.StatusOK, "book_detail.tmpl", gin.H{
				"detail":           data,
				"image":            Image,
				"menu":             menu,
				"chapterList":      chapterList,
				"nextPage":         nextPage,
				"prevPage":         prevPage,
				"hostDefaultImage": HostDefaultImage,
			})
		}
	} else {
		data, _ := index.BookRead(ctx)
		if ctx.GetBool("isMobile") {
			ctx.HTML(http.StatusOK, "m_book_content.tmpl", gin.H{
				"detail":           data,
				"image":            Image,
				"menu":             menu,
				"hostDefaultImage": HostDefaultImage,
			})
		} else {
			ctx.HTML(http.StatusOK, "book_content.tmpl", gin.H{
				"detail":           data,
				"image":            Image,
				"menu":             menu,
				"hostDefaultImage": HostDefaultImage,
			})
		}
	}

}
