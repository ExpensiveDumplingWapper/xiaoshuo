/*
 * @Descripttion: 我见青山多妩媚
 * @Date: 2021-09-29 14:30:50
 * @LastEditTime: 2022-01-10 15:06:50
 */
package routers

import (
	"net/http"

	"xiaoshuo/internal/controllers"
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
	router.Use(middlewares.Debug())
	router.Use(middlewares.Runtime())
	router.ForwardedByClientIP = true

	// 开启 gzip
	router.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{"/metrics"})))

	//监控
	// router.Use(ginprom.PromMiddleware(&ginprom.PromOpts{
	// 	ExcludeRegexStatus:   "404",
	// 	ExcludeRegexEndpoint: exceptPathPattern, //监控需要排除的路由
	// 	EndpointLabelMappingFn: func(c *gin.Context) string {
	// 		return endpointLabelMappingFn(c)
	// 	},
	// }))

	router.NoMethod(HandleNotFound)
	router.NoRoute(HandleNotFound)

	router.LoadHTMLGlob("templates/*")
	router.Static("./static", "static")
	router.GET("/healthz", healthz.Healthz)
	// router.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))
	// router.GET("/", hello)
	router.GET("/", index.Index)

	router.GET("/hello", hellos)
	router.GET("/favicon.ico", hello)
	router.GET("/book_detail/:id", index.BookInfo)
	router.GET("/book_read/:bookid/:chapterid", index.BookRead)
	router.GET("/book_type/category/:booktype/:paga", index.BookType)

	// router.Use(middlewares.Log())
	// 开启 Recover
	// router.Use(middlewares.RecoveryWithZap())

	return router
}

//路由map 监控使用
var maps = []string{
	// "^/oversea/splash/$",
	// "^/oversea/list_banner/(recommend|nearby|online)/$",
	// "^/material/list_banner$",
	exceptPathPattern,
}

func HandleNotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, controllers.NewFailResponse(controllers.ErrNotFound))
	return
}

func hellos(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "indexs.tmpl", gin.H{
		"title": "Main website",
	})
	// ctx.JSON(http.StatusOK, "Hello")
}
func hello(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, "Hello")
}
