/*
 * @Descripttion: 我见青山多妩媚
 * @Date: 2022-01-04 13:56:27
 * @LastEditTime: 2022-01-10 15:53:36
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
type BookInfoDataRes struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    BookInfoData `json:"data"`
}
type BookContentDataRes struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    BookContentData `json:"data"`
}
type BookTypeDataRes struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []BookInfo `json:"data"`
}
type BookContentData struct {
	BookInfo       BookInfo `json:"book"`
	Chapterid      string   `json:"chapterid"`
	Chaptername    string   `json:"chaptername"`
	Chaptercontent string   `json:"chaptercontent"`
	Time           string   `json:"time"`
	Order          int64    `json:"order"`
	Next           string   `json:"next,omitempty"`
	Prev           string   `json:"prev,omitempty"`
}
type BookInfoData struct {
	BookInfo  BookInfo   `json:"bookInfo"`
	Recommend []BookInfo `json:"recommend"`
}
type Recommend struct {
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
	Author        string        `json:"author"`
	Desc          string        `json:"desc,omitempty"`
	Id            string        `json:"id"`
	Name          string        `json:"name"`
	Lastupdate    string        `json:"lastupdate,omitempty"`
	Cover         string        `json:"cover,omitempty"`
	Read          int64         `json:"read,omitempty"`
	BookType      BookType      `json:"type,omitempty"`
	LatestChapter LatestChapter `json:"latestChapter,omitempty"`
	ChapterList   []ChapterList `json:"chapterList,omitempty"`
}
type BookType struct {
	Text string `json:"text"`
	Path string `json:"path"`
}

type LatestChapter struct {
	Chapterid   string `json:"chapterid"`
	Chaptername string `json:"chaptername"`
}

type ChapterList struct {
	Chapterid   string `json:"chapterid"`
	Chaptername string `json:"chaptername"`
}

//首页数据
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

// 小说详情页
func GetBookeInfo(id string) (res BookInfoData) {
	url := "http://180.76.238.148:9093/getBookInfoV2?bookId="
	client := &http.Client{
		// Timeout:   readTimeout,
	}
	url = url + id
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
	var resData BookInfoDataRes
	err = json.Unmarshal(body, &resData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// fmt.Println(resData.Data)
	return resData.Data
}

// 小说内容返回
func GetBookeRead(bookId, chapterId string) (res BookContentData) {
	// url := "http://180.76.238.148:9093/getChapterDetail?bookId=480589&chapterId=2455467"
	url := "http://180.76.238.148:9093/getChapterDetail?bookId="
	client := &http.Client{
		// Timeout:   readTimeout,
	}
	url = url + bookId + "&chapterId=" + chapterId
	fmt.Println(url)
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
	var resData BookContentDataRes
	err = json.Unmarshal(body, &resData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return resData.Data
}

// 小说类型分页
func GetBookeType(bookType, page string) (res []BookInfo) {
	// url := "http://180.76.238.148:9093/getBooks?type=xuanhuanqihuan&size=10&page=1"
	url := "http://180.76.238.148:9093/getBooks?type=" + bookType + "&size=25&page=" + page
	client := &http.Client{
		// Timeout:   readTimeout,
	}
	// url = url + bookId + "&chapterId=" + chapterId
	fmt.Println(url)
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
	var resData BookTypeDataRes
	err = json.Unmarshal(body, &resData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return resData.Data
}

// 小说类型 热门小说返回
func GetBookeTypeHot(bookType string) (res []BookInfo) {
	// url := "http://180.76.238.148:9093/getTopBooksByType?type=xuanhuanqihuan"
	url := "http://180.76.238.148:9093/getTopBooksByType?type=" + bookType
	client := &http.Client{
		// Timeout:   readTimeout,
	}
	fmt.Println(url)
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
	var resData BookTypeDataRes
	err = json.Unmarshal(body, &resData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return resData.Data
}
