package index

import (
	"net/http"
	"xiaoshuo/pkg/log/logrus"

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
	ctx.Redirect(http.StatusMovedPermanently, "http://www.uzwx.com/")
}
