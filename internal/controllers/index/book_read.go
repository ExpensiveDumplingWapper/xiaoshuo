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
