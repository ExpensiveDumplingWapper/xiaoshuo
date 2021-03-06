package routers

import (
	"net/http"

	"xiaoshuo/internal/controllers/healthz"
	"xiaoshuo/internal/controllers/index"
	"xiaoshuo/internal/routers/middlewares"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

//监控里需要排除的路由
const exceptPathPattern = "^(/metrics|/healthz|/|/favicon.ico)$"

func InitRouter() *gin.Engine {

	router := gin.New()
	// router.Use(middlewares.Debug())
	router.Use(middlewares.Runtime())
	router.Use(middlewares.ParseUa())
	router.Use(middlewares.RecoveryWithZap()) // 开启 Recover
	router.ForwardedByClientIP = true

	// 开启 gzip
	router.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{"/metrics"})))

	// router.NoMethod(HandleNotFound)
	// router.NoRoute(HandleNotFound)

	router.NoMethod(index.Index)
	router.NoRoute(index.Index)

	router.LoadHTMLGlob("templates/*")
	router.Static("./static", "static")
	router.GET("/healthz", healthz.Healthz)
	router.GET("/", index.Index)

	router.GET("/hello", hellos)
	router.GET("/favicon.ico", hello)
	router.GET("/book_detail/:id/:page", index.BookInfo)
	router.GET("/category/:booktype/:page", index.BookType)
	router.GET("/book_search", index.BookSearch)
	router.GET("/search_author", index.SearchAuthor)
	router.POST("/leave_message", index.LeavMessage)
	router.POST("/ask_book", index.AskBook)

	router.GET("/leave_message_html", index.LeavMessageHtml)
	router.GET("/ask_book_html", index.AskBookHtml)

	return router
}

//路由map 监控使用
var maps = []string{
	exceptPathPattern,
}

func HandleNotFound(ctx *gin.Context) {
	ctx.Redirect(http.StatusMovedPermanently, "http://www.uzwx.com/")
}

func hellos(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "indexs.tmpl", gin.H{
		"title": "Main website",
	})
}
func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello")
}
