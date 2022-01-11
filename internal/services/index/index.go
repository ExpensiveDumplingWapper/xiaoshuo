/*
 * @Descripttion: 我见青山多妩媚
 * @Date: 2022-01-04 14:41:24
 * @LastEditTime: 2022-01-11 16:59:03
 */
package index

import (
	"xiaoshuo/internal/services/raw_data"

	"github.com/gin-gonic/gin"
)

type HtmlData struct {
	Data1 string `json:"data1"`
	Data2 string `json:"data2"`
	Data3 string `json:"data3"`
	Data4 string `json:"data4"`
	Data5 string `json:"data5"`
}

func Index(ctx *gin.Context) (indexData raw_data.IndexData, hotcontent []raw_data.BookInfo, err error) {
	// data = make(map[string][]HtmlData)

	indexData = raw_data.GetindexData()

	//任取 4个 BookInfo
	hotcontent = make([]raw_data.BookInfo, 4)
	if len(indexData.TopRecommendBooks) > 4 {
		hotcontent = indexData.TopRecommendBooks[:4]
	}

	// fmt.Println(indexData)
	// fmt.Println(data)
	return
}
