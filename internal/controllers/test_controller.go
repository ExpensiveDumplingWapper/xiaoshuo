package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"xiaoshuo/internal/models"
	"xiaoshuo/pkg/log/logger"

	"github.com/gin-gonic/gin"
)

type Test struct {
	ControllerBase
}

type Article struct {
	ArticleId int64  `json:"article_id,omitempty"`
	Info      string `json:"info,omitempty"`
}

func (t *Test) Index(c *gin.Context) {
	//test log
	logger.Log.Info("This is a test for info level.")
	logger.Log.Debug("This a test for debug level.")

	//test zconf
	// data, err := bootstrap.ZConf.GetMysqlConf("/blued/backend/udb/pay_oversea")
	// fmt.Println("%v\n", data)

	//test redis
	client, err := models.RedisCli.GetClient("live")

	if err != nil {
		panic(err)
	}

	client.Set(context.Background(), "foo", "this is a test.", 100*time.Second).Err()
	fmt.Println(client.Get(context.Background(), "foo"))

	////test db
	article := Article{}
	models.BluedDB.Table("article").First(&article)
	fmt.Println(article.Info)

	c.JSON(http.StatusOK, "This is a index for test")

	////test redis
	//redis.GetClient().Set("foo-test", "nihao", 100 * time.Second)
	//result, err := redis.GetClient().Get("foo-test").Result()
	//
	//if err != nil{
	//	Log.Fatalf(err.Error())
	//}
	//
	//Log.Info("Redis test val is %s ", result)
	//
}
