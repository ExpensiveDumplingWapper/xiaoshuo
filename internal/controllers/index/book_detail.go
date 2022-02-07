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

	"xiaoshuo/internal/services/index"
	"xiaoshuo/internal/services/raw_data"

	"github.com/gin-gonic/gin"
)

type ChapterList struct {
	Page string
	Text string
}

func BookInfo(ctx *gin.Context) {

	data, _ := index.BookInfo(ctx)
	menu := raw_data.GetMenus()
	host := ctx.ClientIP() + ":9999"

	chapterList := make([]ChapterList, 0)
	for i := 0; i < data.TotalPage; i++ {
		chapterList = append(chapterList, ChapterList{Page: strconv.Itoa(i + 1), Text: fmt.Sprintf("第%s章-第%s章", strconv.Itoa(1+i*30), strconv.Itoa(30+i*30))})
	}
	// fmt.Println(chapterList)
	// os.Exit(1)
	page := ctx.Param("page")
	finalPage, _ := strconv.Atoi(page)
	nextPage := strconv.Itoa(finalPage + 1)
	prevPageTmp := finalPage - 1
	prevPage := "1"
	if prevPageTmp > 1 {
		prevPage = strconv.Itoa(prevPageTmp)
	}
	if ctx.GetBool("isMobile") {
		ctx.HTML(http.StatusOK, "m_book_detail.tmpl", gin.H{
			"detail":      data,
			"hostServer":  host,
			"image":       Image,
			"menu":        menu,
			"chapterList": chapterList,
			"nextPage":    nextPage,
			"prevPage":    prevPage,
		})
	} else {
		ctx.HTML(http.StatusOK, "book_detail.tmpl", gin.H{
			"detail":      data,
			"hostServer":  host,
			"image":       Image,
			"menu":        menu,
			"chapterList": chapterList,
			"nextPage":    nextPage,
			"prevPage":    prevPage,
		})
	}

}
