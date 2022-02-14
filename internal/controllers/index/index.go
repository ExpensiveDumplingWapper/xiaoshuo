package index

import (
	"net/http"

	"xiaoshuo/internal/services/index"
	"xiaoshuo/internal/services/raw_data"

	"github.com/gin-gonic/gin"
)

var Image = "http://180.76.238.148:8078/image"

const HostDefaultImage = "http://180.76.238.148:8078/image/cover.jpeg"

func Index(ctx *gin.Context) {

	indexData, hotcontent, _ := index.Index(ctx)

	menu := raw_data.GetMenus()
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
			"menu":             menu,
			"indexData":        datas,
			"hostDefaultImage": HostDefaultImage,
			"hotcontent":       hotcontent,
			"image":            Image,
		})
	} else {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"menu":             menu,
			"indexData":        indexData,
			"hostDefaultImage": HostDefaultImage,
			"hotcontent":       hotcontent,
			"image":            Image,
		})
	}

}
