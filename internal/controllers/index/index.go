/*
 * @Descripttion: 我见青山多妩媚
 * @Date: 2022-01-04 14:32:34
 * @LastEditTime: 2022-01-04 19:34:06
 */
package index

import (
	"net/http"

	"xiaoshuo/internal/services/index"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {

	data, indexData, hotcontent, _ := index.Index(ctx)
	// if err != nil {
	// 	ctx.JSON(http.StatusOK, controllers.NewErrResponse(err.Error()))
	// 	return
	// }
	// ctx.JSON(http.StatusOK, controllers.NewSucResponse(data))

	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"menu":       data["menu"],
		"indexData":  indexData,
		"hotcontent": hotcontent,
	})

}
