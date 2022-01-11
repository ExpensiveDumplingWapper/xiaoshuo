/*
 * @Descripttion: 我见青山多妩媚
 * @Date: 2022-01-04 14:32:34
 * @LastEditTime: 2022-01-11 16:47:26
 */
package index

import (
	"fmt"
	"net/http"

	"xiaoshuo/internal/services/index"
	"xiaoshuo/internal/services/raw_data"

	"github.com/gin-gonic/gin"
)

var Image = "http://180.76.238.148:8078/image"

func Index(ctx *gin.Context) {

	indexData, hotcontent, _ := index.Index(ctx)

	menu := raw_data.GetMenus()
	// if err != nil {
	// 	ctx.JSON(http.StatusOK, controllers.NewErrResponse(err.Error()))
	// 	return
	// }
	// ctx.JSON(http.StatusOK, controllers.NewSucResponse(data))

	fmt.Println(ctx.GetBool("isMobile"), 999999)
	if ctx.GetBool("isMobile") {

		var datas [][]raw_data.BookInfo

		datas = append(datas, indexData.BooksByType.Xuanhuanqihuan)
		datas = append(datas, indexData.BooksByType.Xiandaiyanqing)
		datas = append(datas, indexData.BooksByType.Kehuanlingyi)
		datas = append(datas, indexData.BooksByType.Dongfangxuanhuan)
		datas = append(datas, indexData.BooksByType.Wangyoujingji)
		datas = append(datas, indexData.BooksByType.Wuxiaxianxia)
		datas = append(datas, indexData.BooksByType.Xiaoshuotongren)
		datas = append(datas, indexData.BooksByType.Nushengpindao)
		datas = append(datas, indexData.BooksByType.Dushiyanqing)
		datas = append(datas, indexData.BooksByType.Lishijunshi)

		ctx.HTML(http.StatusOK, "m_index.tmpl", gin.H{
			"menu":       menu,
			"indexData":  datas,
			"hotcontent": hotcontent,
			"image":      Image,
		})
	} else {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"menu":       menu,
			"indexData":  indexData,
			"hotcontent": hotcontent,
			"image":      Image,
		})
	}

}
