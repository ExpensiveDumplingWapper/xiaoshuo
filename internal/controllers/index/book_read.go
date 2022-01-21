/*
 * @Descripttion: 我见青山多妩媚
 * @Date: 2022-01-04 14:32:34
 * @LastEditTime: 2022-01-10 11:19:08
 */
package index

import (
	"net/http"

	"xiaoshuo/internal/services/index"
	"xiaoshuo/internal/services/raw_data"

	"github.com/gin-gonic/gin"
)

func BookRead(ctx *gin.Context) {

	data, _ := index.BookRead(ctx)
	menu := raw_data.GetMenus()
	if ctx.GetBool("isMobile") {
		ctx.HTML(http.StatusOK, "m_book_content.tmpl", gin.H{
			"detail": data,
			"image":  Image,
			"menu":   menu,
		})
	} else {
		ctx.HTML(http.StatusOK, "book_content.tmpl", gin.H{
			"detail": data,
			"image":  Image,
			"menu":   menu,
		})
	}

}
