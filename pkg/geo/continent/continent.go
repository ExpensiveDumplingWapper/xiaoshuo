package continent

import (
	"regexp"
)

var codes = map[string]map[string]string{
	"0": {"code": "0", "alpha2": "UN"},
	"1": {"code": "1", "alpha2": "AS"},
	"2": {"code": "2", "alpha2": "EU"},
	"3": {"code": "3", "alpha2": "AF"},
	"4": {"code": "4", "alpha2": "OC"},
	"5": {"code": "5", "alpha2": "NA"},
	"6": {"code": "6", "alpha2": "SA"},
	"7": {"code": "7", "alpha2": "AN"},
}

var names = map[string]string{
	"UN": "0",
	"AS": "1",
	"AP": "1",
	"EU": "2",
	"AF": "3",
	"OC": "4",
	"OA": "4",
	"NA": "5",
	"SA": "6",
	"LA": "6",
	"AN": "7",
	"AQ": "7",
}

//var names = map[string]interface{}{
//	"0": map[string]string{"en": "World", "zh_cn": "世界", "zh_tw": "世界"},
//	"1": map[string]string{"en": "Asia", "zh_cn": "亚洲", "zh_tw": "亞洲"},
//	"2": map[string]string{"en": "Europe", "zh_cn": "欧洲", "zh_tw": "歐洲"},
//	"3": map[string]string{"en": "Africa", "zh_cn": "非洲", "zh_tw": "非洲"},
//	"4": map[string]string{"en": "Oceania", "zh_cn": "大洋洲", "zh_tw": "大洋洲"},
//	"5": map[string]string{"en": "North America", "zh_cn": "北美洲", "zh_tw": "北美洲"},
//	"6": map[string]string{"en": "Latin America", "zh_cn": "拉丁美洲", "zh_tw": "拉丁美洲"},
//	"7": map[string]string{"en": "Antarctica", "zh_cn": "南极洲", "zh_tw": "南極洲"},
//}

func Get(k string) map[string]string {
	var res = map[string]string{}

	pattern := "\\d+"
	isNum, _ := regexp.MatchString(pattern, k)
	if isNum {
		if r, ok := codes[k]; ok {
			return r
		}
	} else {
		if num, ok := names[k]; ok {
			if r, ok2 := codes[num]; ok2 {
				return r
			}
		}
	}

	return res
}
