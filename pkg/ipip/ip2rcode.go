package ipip

import (
	"fmt"
	"net"
	geoContinent "xiaoshuo/pkg/geo/continent"
	geoCountry "xiaoshuo/pkg/geo/country"

	"github.com/ipipdotnet/ipdb-go"
)

var chinaCode = map[string]string{
	"TW": "156",
	"HK": "156",
	"MO": "156",
}

func IP2Rcode(ip string, ipv4db *ipdb.City, ipv6db *ipdb.City) string {
	if ip == "" {
		return "4_036_000000"
	}
	db := ipv4db

	netIP := net.ParseIP(ip)
	if netIP.To4() == nil {
		db = ipv6db
	}

	resMap, _ := db.FindMap(ip, "CN")

	chinaAdminCode := "000000"
	if _, ok := resMap["china_admin_code"]; ok && len(resMap["china_admin_code"]) > 0 {
		chinaAdminCode = resMap["china_admin_code"]
	}

	countryName := "000"
	if _, ok := resMap["country_code"]; ok && len(resMap["country_code"]) > 0 {
		countryName = resMap["country_code"]
	}

	continentName := "0"
	if _, ok := resMap["continent_code"]; ok && len(resMap["continent_code"]) > 0 {
		continentName = resMap["continent_code"]
	}

	if chinaAdminCode == "000000" && continentName == "0" && countryName == "000" {
		return "0_000_000000"
	}

	countryCode := "000"
	countryInfo := map[string]string{}

	if countryName != "000" {
		//取国家信息
		countryInfo = geoCountry.Get(countryName)
		//取国家码
		if _, hasCode := countryInfo["code"]; hasCode {
			countryCode = countryInfo["code"]
		}
	}

	continentCode := "0"
	if continentName != "0" {
		//取大洲码
		//先从国家信息里取大洲码
		if _, hasContinent := countryInfo["continent"]; hasContinent {
			continentCode = countryInfo["continent"]
		} else {
			//如果国家里没有大洲码 再取大洲
			continentInfo := geoContinent.Get(continentName)
			if _, hasCode := continentInfo["code"]; hasCode {
				continentCode = continentInfo["code"]
			}
		}
	}

	//香港、台湾、澳门
	if _, is := chinaCode[countryName]; is {
		countryCode = "156"
	}

	region := fmt.Sprintf("%s_%s_%s", continentCode, countryCode, chinaAdminCode)

	fmt.Println(ip, region)

	return region
}
