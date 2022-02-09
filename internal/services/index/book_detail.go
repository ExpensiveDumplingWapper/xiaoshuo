package index

import (
	"strings"
	"xiaoshuo/internal/services/raw_data"

	"github.com/gin-gonic/gin"
)

func BookInfo(ctx *gin.Context) (res raw_data.BookChaptersData, err error) {
	id := ctx.Param("id")
	page := strings.ReplaceAll(ctx.Param("page"), "c_", "")

	if page == "" {
		page = "1"
	}
	res = raw_data.GetBookeInfo(id, page)
	return
}
