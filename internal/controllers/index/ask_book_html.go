package index

import (
	"net/http"
	"xiaoshuo/internal/services/raw_data"

	"github.com/gin-gonic/gin"
)

func AskBookHtml(ctx *gin.Context) {

	menu := raw_data.GetMenus()
	ctx.HTML(http.StatusOK, "ask_book.tmpl", gin.H{
		"menu":             menu,
		"hostDefaultImage": HostDefaultImage,
		"image":            Image,
	})
}
