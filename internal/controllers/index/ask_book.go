package index

import (
	"net/http"
	"xiaoshuo/pkg/log/logrus"

	"github.com/gin-gonic/gin"
)

func AskBook(ctx *gin.Context) {
	email := ctx.PostForm("email")
	message := ctx.PostForm("message")
	if email != "" || message != "" {
		logrus.AskBookDB(email, message, ctx.ClientIP())
	}
	ctx.Redirect(http.StatusMovedPermanently, "http://www.uzwx.com/")
}
