/*
 * @Descripttion: 我见青山多妩媚
 * @Date: 2022-01-04 13:56:27
 * @LastEditTime: 2022-01-06 18:36:38
 */
package raw_data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MenusRes struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    []Menus `json:"data"`
}

type Menus struct {
	Text string `json:"text"`
	Path string `json:"path"`
}

func GetMenus() (res []Menus) {
	url := "http://180.76.238.148:9093/getMenus"
	client := &http.Client{
		// Timeout:   readTimeout,
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var resData MenusRes
	err = json.Unmarshal(body, &resData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return resData.Data
}
