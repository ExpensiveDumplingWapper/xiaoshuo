/*
 * @Descripttion: 我见青山多妩媚
 * @Date: 2022-01-04 13:56:27
 * @LastEditTime: 2022-01-04 22:33:58
 */
package raw_data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IndexDataRes struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    IndexData `json:"data"`
}
type IndexData struct {
	BooksByType       BooksByType `json:"booksByType"`
	TopRecommendBooks []BookInfo  `json:"topRecommendBooks"`
	LastupdateBooks   []BookInfo  `json:"lastupdateBooks"`
}
type BooksByType struct {
	Xuanhuanqihuan   []BookInfo `json:"xuanhuanqihuan,omitempty"`
	Xiandaiyanqing   []BookInfo `json:"xiandaiyanqing,omitempty"`
	Kehuanlingyi     []BookInfo `json:"kehuanlingyi,omitempty"`
	Dongfangxuanhuan []BookInfo `json:"dongfangxuanhuan,omitempty"`
	Wangyoujingji    []BookInfo `json:"wangyoujingji,omitempty"`
	Wuxiaxianxia     []BookInfo `json:"wuxiaxianxia,omitempty"`
	Xiaoshuotongren  []BookInfo `json:"xiaoshuotongren,omitempty"`
	Nushengpindao    []BookInfo `json:"nu:shengpindao,omitempty"`
	Dushiyanqing     []BookInfo `json:"dushiyanqing,omitempty"`
	Lishijunshi      []BookInfo `json:"lishijunshi,omitempty"`
}

type BookInfo struct {
	Author     string   `json:"author"`
	Desc       string   `json:"desc"`
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Lastupdate string   `json:"lastupdate"`
	Cover      string   `json:"cover"`
	Read       int64    `json:"read"`
	BookType   BookType `json:"type"`
}
type BookType struct {
	Text string `json:"text"`
	Path string `json:"path"`
}

func GetindexData() (res IndexData) {
	url := "http://180.76.238.148:9093/indexData"
	client := &http.Client{
		// Timeout:   readTimeout,
	}
	// bytesData, _ := json.Marshal(data)
	// req, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	// req.Header.Add("Content-Type", "application/json")
	// resp, err := client.Do(req)
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var resData IndexDataRes
	err = json.Unmarshal(body, &resData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// fmt.Println(resData.Data)
	return resData.Data
}
