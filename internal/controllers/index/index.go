/*
 * @Descripttion: 我见青山多妩媚
 * @Date: 2022-01-04 14:32:34
 * @LastEditTime: 2022-01-12 10:50:36
 */
package index

import (
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

	// fmt.Println(ctx.GetBool("isMobile"), 999999)
	if ctx.GetBool("isMobile") {

		var datas [][]raw_data.BookInfo

		if len(indexData.BooksByType.Xuanhuanqihuan) > 0 {
			datas = append(datas, indexData.BooksByType.Xuanhuanqihuan)
		}

		if len(indexData.BooksByType.Xiandaiyanqing) > 0 {
			datas = append(datas, indexData.BooksByType.Xiandaiyanqing)
		}
		if len(indexData.BooksByType.Kehuanlingyi) > 0 {
			datas = append(datas, indexData.BooksByType.Kehuanlingyi)
		}
		if len(indexData.BooksByType.Dongfangxuanhuan) > 0 {
			datas = append(datas, indexData.BooksByType.Dongfangxuanhuan)
		}
		if len(indexData.BooksByType.Wangyoujingji) > 0 {
			datas = append(datas, indexData.BooksByType.Wangyoujingji)
		}

		if len(indexData.BooksByType.Wuxiaxianxia) > 0 {
			datas = append(datas, indexData.BooksByType.Wuxiaxianxia)
		}
		if len(indexData.BooksByType.Xiaoshuotongren) > 0 {
			datas = append(datas, indexData.BooksByType.Xiaoshuotongren)
		}
		if len(indexData.BooksByType.Nushengpindao) > 0 {
			datas = append(datas, indexData.BooksByType.Nushengpindao)
		}
		if len(indexData.BooksByType.Dushiyanqing) > 0 {
			datas = append(datas, indexData.BooksByType.Dushiyanqing)
		}
		if len(indexData.BooksByType.Lishijunshi) > 0 {
			datas = append(datas, indexData.BooksByType.Lishijunshi)
		}
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
