/*
 * @Descripttion: 我见青山多妩媚
 * @Date: 2022-01-04 14:32:34
 * @LastEditTime: 2022-01-11 18:32:44
 */
package index

import (
	"net/http"

	"xiaoshuo/internal/services/index"
	"xiaoshuo/internal/services/raw_data"

	"github.com/gin-gonic/gin"
)

func BookInfo(ctx *gin.Context) {

	data, _ := index.BookInfo(ctx)
	menu := raw_data.GetMenus()
	host := ctx.ClientIP() + ":9999"
	// if err != nil {
	// 	ctx.JSON(http.StatusOK, controllers.NewErrResponse(err.Error()))
	// 	return
	// }
	// ctx.JSON(http.StatusOK, controllers.NewSucResponse(data))

	if ctx.GetBool("isMobile") {
		ctx.HTML(http.StatusOK, "m_book_detail.tmpl", gin.H{
			"detail":     data,
			"hostServer": host,
			"image":      Image,
			"menu":       menu,
		})
	} else {
		ctx.HTML(http.StatusOK, "book_detail.tmpl", gin.H{
			"detail":     data,
			"hostServer": host,
			"image":      Image,
			"menu":       menu,
		})
	}

}
