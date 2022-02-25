package raw_data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

type BookChapterRes struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    BookChaptersData `json:"data"`
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

type BookChaptersData struct {
	BookInfo  BookInfo   `json:"bookInfo"`
	Chapters  []Chapter  `json:"chapters"`
	TotalPage int        `json:"total_page"`
	Recommend []BookInfo `json:"recommend"`
}
type Chapter struct {
	Chapterid      string `json:"chapterid"`
	Chaptername    string `json:"chaptername"`
	Chaptercontent string `json:"chaptercontent,omitempty"`
	Time           string `json:"time,omitempty"`
	Order          int    `json:"order,omitempty"`
	Book           string `json:"book,omitempty"`
	Next           string `json:"next,omitempty"`
	Prev           string `json:"prev,omitempty"`
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

const host = "http://180.76.238.148:9093"

//首页数据
func GetindexData(ctx *gin.Context) (res IndexData) {
	url := host + "/indexData"
	client := &http.Client{
		// Timeout:   readTimeout,
	}
	fmt.Println(ctx.ClientIP(), time.Now().Format("2006-01-02 15:04:05"), url)
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
	return resData.Data
}

// 小说详情页
func GetBookeInfo(ctx *gin.Context, id, page string) (res BookChaptersData) {
	url := host + "/getChapters?bookId=" + id + "&page=" + page + "&size=30"
	client := &http.Client{
		// Timeout:   readTimeout,
	}
	fmt.Println(ctx.ClientIP(), time.Now().Format("2006-01-02 15:04:05"), url)
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
	var resData BookChapterRes
	err = json.Unmarshal(body, &resData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return resData.Data
}

// 小说内容返回
func GetBookeRead(ctx *gin.Context, bookId, chapterId string) (res BookContentData) {
	url := host + "/getChapterDetail?bookId="
	client := &http.Client{
		// Timeout:   readTimeout,
	}
	url = url + bookId + "&chapterId=" + chapterId
	fmt.Println(ctx.ClientIP(), time.Now().Format("2006-01-02 15:04:05"), url)
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
func GetBookeType(ctx *gin.Context, bookType, page string) (res []BookInfo) {
	url := host + "/getBooks?type=" + bookType + "&size=25&page=" + page
	client := &http.Client{
		// Timeout:   readTimeout,
	}
	fmt.Println(ctx.ClientIP(), time.Now().Format("2006-01-02 15:04:05"), url)
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
func GetBookeTypeHot(ctx *gin.Context, bookType string) (res []BookInfo) {
	url := host + "/getTopBooksByType?type=" + bookType
	client := &http.Client{
		// Timeout:   readTimeout,
	}
	fmt.Println(ctx.ClientIP(), time.Now().Format("2006-01-02 15:04:05"), url)
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

// 小说查询 作者和书名 然后合并返回
func GetBookSearch(ctx *gin.Context, keyWord, page string) (res []BookInfo) {
	urlAuthor := host + "/search?size=30&author=" + keyWord + "&page=" + page
	urlName := host + "/search?size=30&name=" + keyWord + "&page=" + page

	urlAuthorData := GetBookSearchData(ctx, urlAuthor)
	urlurlName := GetBookSearchData(ctx, urlName)
	res = append(res, urlurlName...)
	res = append(res, urlAuthorData...)
	return
}

// 小说查询 作者
func GetSearchAuthor(ctx *gin.Context, keyWord, page string) (res []BookInfo) {
	urlAuthor := host + "/search?size=30&author=" + keyWord + "&page=" + page
	urlAuthorData := GetBookSearchData(ctx, urlAuthor)
	res = append(res, urlAuthorData...)
	return
}

func GetBookSearchData(ctx *gin.Context, url string) (res []BookInfo) {
	client := &http.Client{
		// Timeout:   readTimeout,
	}
	fmt.Println(ctx.ClientIP(), time.Now().Format("2006-01-02 15:04:05"), url)
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
