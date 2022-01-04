package user

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"xiaoshuo/pkg/log/logger"

	HTTPClient "git.ourbluecity.com/overseas-server-tools/go-util-http"
	"github.com/gin-gonic/gin"
)

type UserResult struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	DATA UserInfo
}

func GetUsersProfiles(c *gin.Context, target int, field []string, currentInfo, isPassport int) (UserInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	var user = UserInfo{}
	fields := strings.Join(field, ",")
	isPassportTemp := 0
	temp := false
	if strings.Contains(fields, "smid") || strings.Contains(fields, "dev_dna") {
		isPassportTemp = 3
		temp = true
	}
	if isPassport > 0 {
		isPassportTemp = isPassport
		temp = true
	}
	queryParam := make(map[string]string)
	if temp {
		queryParam = map[string]string{
			"grant_fields": fields,
			"current_info": strconv.Itoa(currentInfo),
			"is_passport":  strconv.Itoa(isPassportTemp),
		}
	} else {
		queryParam = map[string]string{
			"grant_fields": fields,
			"current_info": strconv.Itoa(currentInfo),
		}
	}
	url := "/users/" + strconv.Itoa(target)
	reslust, err := baseClient.Get(ctx, HTTPTools, HTTPClient.QueryReq{
		URL:        url,
		QueryParam: queryParam,
		HeaderParam: map[string]string{
			"x-iris-uid": "12",
		},
	})
	if err != nil {
		logger.Log.Errorf(" http GetUsersProfiles error:" + err.Error())
		return user, err
	}
	var res UserResult
	if err := json.Unmarshal(reslust.ResBody, &res); err != nil {
		logger.Log.Errorf(" http GetUsersProfiles json:" + err.Error())
		return user, err
	}
	if reslust.StatusCode != 200 || reslust.Status != "200 OK" {
		logger.Log.Errorf(" http GetUsersProfiles error:" + err.Error() + "code:" + strconv.Itoa(reslust.StatusCode) + " status:" + reslust.Status + "result:" + strconv.Itoa(res.Code) + res.Msg)
		return user, err
	}
	return res.DATA, err
}

type UserInfo struct {
	Uid     int    `json:"uid"`
	Version string `json:"version"`
}
