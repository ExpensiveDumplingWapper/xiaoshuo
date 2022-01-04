package healthz

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthz(c *gin.Context) {
	//redisClient, err := models.RedisCli.GetClient("adms_read")
	//if err != nil {
	//	fmt.Println("get redis instance error:"+err.Error())
	//	c.Writer.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//
	//result, e := redisClient.Ping(c).Result()
	//if e != nil {
	//	fmt.Println("ping redis error:"+e.Error())
	//	c.Writer.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//
	//if result != "PONG" {
	//	fmt.Println("ping redis error, expect: PONG, got: "+ result)
	//	c.Writer.WriteHeader(http.StatusInternalServerError)
	//	return
	//}

	c.Writer.WriteHeader(http.StatusOK)
}
