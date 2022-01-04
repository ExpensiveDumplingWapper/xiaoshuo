package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"time"
	log "xiaoshuo/pkg/log/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Info(logName string, description string, ctx *gin.Context) {
	requestEntry(logName, "i", ctx).Info(description)
}

func Error(logName string, description string, ctx *gin.Context) {
	requestEntry(logName, "e", ctx).Error(description)
}

func Debug(logName string, description string, ctx *gin.Context) {
	requestEntry(logName, "d", ctx).Debug(description)
}

func Warn(logName string, description string, ctx *gin.Context) {
	requestEntry(logName, "w", ctx).Warn(description)
}

func Success(returnAids string, ctx *gin.Context) {
	Info("success", returnAids, ctx)
}

func Empty(ctx *gin.Context) {
	Info("empty", "Ads Empty", ctx)
}

func SysError(ctx *gin.Context, description string) {
	if _, file, line, ok := runtime.Caller(1); ok {
		description = fmt.Sprintf("file: %s \nline: %d \n%s", file, line, description)
	}
	Error("sys_error", description, ctx)
}

var headerName = []string{
	"User-Agent",
	"X-Remote-Userid",
	"X-Forwarded-For",
	"X-Real-Client-Ip",
	"X-Real-Ip",
	"X-Guest-Uid",
	"X-Remote-Channel",
	"x-Guest-Uid",
}

func requestEntry(logName string, level string, ctx *gin.Context) *logrus.Entry {
	headers := make(map[string]string)
	for _, name := range headerName {
		if v := ctx.GetHeader(name); len(v) > 0 {
			headers[name] = v
		}
	}
	jsonHeader, _ := json.Marshal(headers)
	hostname, _ := os.Hostname()

	logContent := logrus.Fields{
		"time":           time.Now().UnixNano() / 1e6,
		"service":        "adms-go",
		"log_name":       logName,
		"level":          level,
		"uid":            "",
		"client_ip":      "",
		"server_ip":      hostname,
		"request_id":     "",
		"request_url":    ctx.Request.RequestURI,
		"request_type":   ctx.Request.Method,
		"request_header": string(jsonHeader),
	}

	//Test Only
	//fmt.Println(logContent)

	return log.Log.WithFields(logContent)
}
