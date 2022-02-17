package index

import (
	"net/http"
	"xiaoshuo/pkg/log/logrus"
	"xiaoshuo/pkg/mail"

	"github.com/gin-gonic/gin"
)

func AskBook(ctx *gin.Context) {
	email := ctx.PostForm("email")
	message := ctx.PostForm("message")
	if email != "" || message != "" {
		logrus.AskBookDB(email, message, ctx.ClientIP())
	}
	mail.SendMail([]string{"1071973064@qq.com"}, "求书加书", message)
	ctx.Redirect(http.StatusMovedPermanently, "http://www.uzwx.com/")
}
