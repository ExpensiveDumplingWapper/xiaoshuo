package request

import (
	"fmt"
	"regexp"
	"strings"
	"xiaoshuo/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
	"github.com/speps/go-hashids/v2"
)

type UaStruct struct {
	Device           string
	Platform         string
	PlatformVersion  string
	Screen           [2]int
	VersionName      string
	VersionCode      int
	Timezone         string
	Additions        string
	Region           string
	PlatformCode     int
	App              int
	ScreenResolution string
	ScreenProportion float64
	IDFA             string
	IMEI             string
	MAC              string
	UID              int
	HashUID          string
	GuestUID         int
}

const UaPattern = `(?i)^Mozilla/5\.0 \((.+)\) .*(ios|android|windowsphone|windows|mac)/([^ ]+) \(([^ ]+)\)(.*)$`

func HttpParseUa(ctx *gin.Context) UaStruct {
	ua := UaStruct{
		Device:           "",
		Platform:         "unknown",
		Screen:           [2]int{1080, 1920},
		ScreenResolution: parseScreenResolution("1080", "1920"),
		ScreenProportion: parseScreenProportion("1080", "1920"),
		VersionName:      "",
		VersionCode:      0,
		Timezone:         "Asia/Shanghai",
		Additions:        "",
		Region:           "", //ipip.IP2Rcode(GetRemoteIP(ctx), models.IPV4DB, models.IPV6DB),
		PlatformCode:     0,
		App:              2,
		IDFA:             ctx.GetHeader("x-remote-idfa"),
		IMEI:             ctx.GetHeader("x-remote-imei"),
		MAC:              ctx.GetHeader("x-remote-mac"),
		UID:              GetRemoteUserID(ctx),
		HashUID:          EncodeHashUserID(GetRemoteUserID(ctx)),
		GuestUID:         GetGuestUid(ctx),
	}

	userAgent := ctx.GetHeader("User-Agent")
	r, _ := regexp.Compile(UaPattern)
	match := r.MatchString(userAgent)

	if match == false {
		return ua
	}

	uaData := r.FindStringSubmatch(userAgent)
	deviceInfo := strings.Split(utils.ReverseString(uaData[3]), "_")

	ua.Device = uaData[1]
	ua.Platform = strings.ToLower(uaData[2])
	ua.Screen = [2]int{utils.Str2Int(deviceInfo[0], 1080), utils.Str2Int(deviceInfo[1], 1920)}
	ua.ScreenResolution = parseScreenResolution(deviceInfo[0], deviceInfo[1])
	ua.ScreenProportion = parseScreenProportion(deviceInfo[0], deviceInfo[1])
	ua.VersionName = deviceInfo[2]
	ua.VersionCode = utils.Str2Int(deviceInfo[3], 0)
	ua.Timezone = uaData[4]
	ua.Additions = uaData[5]
	ua.PlatformCode = parsePlatformCode(strings.ToLower(uaData[2]))
	ua.App = utils.Str2Int(strings.ReplaceAll(uaData[5], "app/", ""), 2)
	d := strings.Split(uaData[0], ";")
	for _, i := range d {
		t := strings.ToLower(strings.ReplaceAll(i, " ", ""))
		if strings.Contains(t, ua.Platform) {
			ua.PlatformVersion = strings.ReplaceAll(t, ua.Platform, "")
			break
		}
	}

	return ua
}

func parsePlatformCode(name string) int {
	switch name {
	case "android":
		return 1
	case "ios":
		return 2
	default:
		return 0
	}
}

func parseScreenResolution(width, height string) string {
	sr := fmt.Sprintf("%sx%s", width, height)

	var resolutions = []string{
		"1242x2688", //iphone max
		"750x1624",  //iphone xr
		"828x1792",  //iphone xr
		"1125x2436", // iphone x
	}

	for _, r := range resolutions {
		if r == sr {
			sr = "1125x2436"
		}
	}

	return sr
}

func parseScreenProportion(width, height string) float64 {
	if utils.Str2Int(width, 0) > 0 {
		return float64(utils.Str2Int64(width, 1080)) / float64(utils.Str2Int64(height, 1920))

	}
	return 1080.0 / 1920.0
}

//获取请求ID
func GetRequestID(ctx *gin.Context) string {
	reqId := ctx.GetHeader("X-Request-Id")
	if len(reqId) == 0 {
		reqId, _ = uuid.GenerateUUID()
		reqId = strings.ReplaceAll(reqId, "-", "")
	}
	return reqId
}

//获取用户ID
func GetRemoteUserID(ctx *gin.Context) int {
	uidString := ctx.GetHeader("X-Remote-Userid")
	if len(uidString) == 0 {
		return 0
	}

	match, _ := regexp.MatchString(`^\d+$`, uidString)
	if match == false {
		return DecodeHashUserID(uidString)
	}

	id := utils.Str2Int(uidString, 0)
	return id
}

//获取用户IP
func GetRemoteIP(ctx *gin.Context) string {
	ip := ctx.GetHeader("True-Client-Ip")
	if len(ip) == 0 {
		ip = ctx.GetHeader("X-Real-Ip") //X-Real-Ip
		if len(ip) == 0 {
			ip = ctx.ClientIP()
		}
	}
	return ip
}

//获取VIP免广告标识
func GetVipFilter(ctx *gin.Context) int {
	filter := ctx.DefaultQuery("svip_filter", "0")
	res := utils.Str2Int(filter, 0)

	return res
}

//获取金主免广告标识
func GetRichLevelFilter(ctx *gin.Context) int {
	filter := ctx.DefaultQuery("rich_level_filter", "0")
	res := utils.Str2Int(filter, 0)

	return res
}

//获取新用户免广告标识
func GetNewUserFilter(ctx *gin.Context) int {
	filter := ctx.DefaultQuery("new_user_filter", "0")
	res := utils.Str2Int(filter, 0)

	return res
}

//获取需要排除的广告ID
func GetExcludeID(ctx *gin.Context) []int {
	var ids = []int{}
	idString := ctx.DefaultQuery("exclude_id", "")
	if len(idString) == 0 {
		return ids
	}

	idArr := strings.Split(idString, ",")
	if len(idArr) == 0 {
		return ids
	}

	for _, id := range idArr {
		ids = append(ids, utils.Str2Int(id, 0))
	}

	return ids
}

func EncodeHashUserID(uid int) string {
	hd := hashids.NewData()
	hd.Salt = "1766"
	hd.Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123567890"
	hd.MinLength = 6
	h, _ := hashids.NewWithData(hd)
	d, _ := h.Encode([]int{uid})

	return d
}

func DecodeHashUserID(hid string) int {
	hd := hashids.NewData()
	hd.Salt = "1766"
	hd.Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123567890"
	hd.MinLength = 6
	h, _ := hashids.NewWithData(hd)
	uids, _ := h.DecodeWithError(hid)

	if len(uids) > 0 {
		return uids[0]
	}

	return 0
}

func GetAcceptLanguage(ctx *gin.Context) string {
	lang := ctx.GetHeader("accept-language")

	if len(lang) == 0 {
		lang = "en-us"
	}

	return lang
}

func GetHashUID(ctx *gin.Context) string {
	return EncodeHashUserID(GetRemoteUserID(ctx))
}

func GetGuestUid(ctx *gin.Context) int {
	uid := ctx.GetHeader("x-guest-uid")
	if len(uid) == 0 {
		return 0
	}

	return utils.Str2Int(uid, 0)
}

func UID(ctx *gin.Context) int64 {
	val, ok := ctx.Get("uid")
	if !ok {
		return 0
	}

	return val.(int64)
}
