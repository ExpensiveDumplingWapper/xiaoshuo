package index

import (
	"net/http"
	"xiaoshuo/internal/services/raw_data"

	"github.com/gin-gonic/gin"
)

func LeavMessageHtml(ctx *gin.Context) {

	menu := raw_data.GetMenus()
	ctx.HTML(http.StatusOK, "leave_message.tmpl", gin.H{
		"menu":             menu,
		"hostDefaultImage": HostDefaultImage,
		"image":            Image,
	})
}
