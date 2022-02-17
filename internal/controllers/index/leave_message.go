package index

import (
	"net/http"
	"xiaoshuo/pkg/log/logrus"
	"xiaoshuo/pkg/mail"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Title   string `json:"title,omitempty"`
	Message string `json:"message,omitempty"`
}

func LeavMessage(ctx *gin.Context) {
	email := ctx.PostForm("email")
	message := ctx.PostForm("message")
	if email != "" || message != "" {
		logrus.LeavMessDB(email, message, ctx.ClientIP())
	}
	// ctx.JSON(http.StatusOK, "")

	mail.SendMail([]string{"1071973064@qq.com"}, "留言反馈", message)
	ctx.Redirect(http.StatusMovedPermanently, "http://www.uzwx.com/")
}
