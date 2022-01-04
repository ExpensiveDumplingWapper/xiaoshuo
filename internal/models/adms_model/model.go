package adms_model

import (
	"context"
	"encoding/json"
	"sort"
	"strings"
	"xiaoshuo/internal/utils"

	"github.com/hashicorp/go-version"

	"xiaoshuo/internal/models"

	"gorm.io/gorm"
)

type Adm struct {
	Id           int    `gorm:"id;" json:"id,omitempty"`
	Aaid         int    `gorm:"aaid;" json:"aaid,omitempty"`
	Title        string `gorm:"name;" json:"name,omitempty"`
	Status       int    `gorm:"status;" json:"status,omitempty"`
	Type         int    `gorm:"type;" json:"type,omitempty"`
	Groups       int    `gorm:"groups;" json:"groups,omitempty"`
	Position     int    `gorm:"position;" json:"position,omitempty"`
	Purpose      int    `gorm:"purpose;" json:"purpose,omitempty"`
	Material     string `gorm:"material;" json:"material,omitempty"`
	MaterialType string `gorm:"material_type;" json:"material_type,omitempty"`
	JumpUrl      string `gorm:"jump_url;" json:"jump_url,omitempty"`
	Region       string `gorm:"region;" json:"region,omitempty"`
	Ranking      int    `gorm:"ranking;" json:"ranking,omitempty"`
	Channel      string `gorm:"channel;" json:"channel,omitempty"`
	Avatar       string `gorm:"avatar;" json:"avatar,omitempty"`
	Nickname     string `gorm:"nickname;" json:"nickname,omitempty"`
	Content      string `gorm:"content;" json:"content,omitempty"`
	DetailsUrl   string `gorm:"details_url;" json:"details_url,omitempty"`
	ShowMembers  int    `gorm:"show_members;" json:"show_members,omitempty"`
	CloseTime    int    `gorm:"close_time;" json:"close_time,omitempty"`
	CanClose     int    `gorm:"can_close;" json:"can_close,omitempty"`

	Platform          int    `gorm:"platform;" json:"platform,omitempty"`
	PutType           int    `gorm:"put_type;" json:"put_type,omitempty"`
	StartTime         int    `gorm:"start_time;" json:"StartTime,omitempty"`
	EndTime           int    `gorm:"end_time;" json:"end_time,omitempty"`
	Strategy          string `gorm:"strategy;" json:"strategy,omitempty"`
	AthId             string `gorm:"ath_id;" json:"ath_id,omitempty"`
	SplashTime        int    `gorm:"splash_time;" json:"splash_time,omitempty"`
	SplashForceTime   int    `gorm:"splash_force_time;" json:"splash_force_time,omitempty"`
	FrequencyType     int    `gorm:"frequency_type;" json:"frequency_type,omitempty"`
	FrequencyCount    int    `gorm:"frequency_count;" json:"frequency_count,omitempty"`
	FrequencyInterval int    `gorm:"frequency_interval;" json:"frequency_interval,omitempty"`
	FrequencyDays     int    `gorm:"frequency_days;" json:"frequency_days,omitempty"`
	ClientVersion     string `gorm:"client_version;" json:"client_version,omitempty"`
	PutTableName      string `gorm:"put_table_name;" json:"put_table_name,omitempty"`
	SellType          int    `gorm:"sell_type;" json:"sell_type,omitempty"`
	InspiringTime     int    `gorm:"inspiring_time;" json:"inspiring_time,omitempty"`
	ClickNum          int    `gorm:"click_num;" json:"click_num,omitempty"`
	InspireWatchNum   int    `gorm:"inspire_watch_num;" json:"inspire_watch_num,omitempty"`
	IntervalNum       int    `gorm:"interval_num;" json:"interval_num,omitempty"`
	InsertRow         int    `gorm:"insert_row;" json:"insert_row,omitempty"`
}

type ClientVersion struct {
	CheckType    int    `json:"checkType"`
	StartVersion string `json:"startVersion"`
	EndVersion   string `json:"endVersion"`
}

type Strategy struct {
	Mantissa []int `json:"mantissa"`
}

func (adm *Adm) TableName() string {
	return "adm_oversea"
}

func List(ctx context.Context) []Adm {
	admsData := Adm{
		Status: 1,
	}
	err, items := CheckByRoomId(ctx, models.BluedDB, &admsData)
	if err != nil {
		return []Adm{}
	}
	return items
}

func (adm *Adm) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         adm.Id,
		"title":      adm.Title,
		"status":     adm.Status,
		"type":       adm.Type,
		"groups":     adm.Groups,
		"position":   adm.Position,
		"purpose":    adm.Purpose,
		"platform":   adm.Platform,
		"put_type":   adm.PutType,
		"start_time": adm.StartTime,
		"end_time":   adm.EndTime,
		"strategy":   adm.Strategy,
		"ath_id":     adm.AthId,
	}
}

func CheckByRoomId(ctx context.Context, tx *gorm.DB, admsData *Adm) (err error, res []Adm) {
	tx.Model(&admsData).Where("status=?", admsData.Status).Order("id asc").Find(&res)
	return nil, res
}

type CacheAdm struct {
	Id                int
	Aaid              int
	Title             string
	Status            int
	Type              int
	Groups            int
	Position          int
	Purpose           int
	Material          string
	MaterialType      string
	JumpUrl           string
	Region            *[]RegionsMap
	Ranking           int
	Channel           string
	Avatar            string
	Nickname          string
	Content           string
	DetailsUrl        string
	ShowMembers       int
	CloseTime         int
	CanClose          int
	Platform          int
	PutType           int
	StartTime         int
	EndTime           int
	Strategy          Strategy
	AthId             string
	SplashTime        int
	SplashForceTime   int
	FrequencyType     int
	FrequencyCount    int
	FrequencyInterval int
	FrequencyDays     int
	ClientVersion     CacheClientVersion
	PutTableName      string
	SellType          int
	InspiringTime     int
	ClickNum          int
	InspireWatchNum   int
	IntervalNum       int
	InsertRow         int
}

func GetAllAdms(ctx context.Context) []CacheAdm {
	return formatAdms(List(ctx))
}

//格式化广告
func formatAdms(adms []Adm) []CacheAdm {
	cAdms := make([]CacheAdm, len(adms))

	for k, adm := range adms {
		cAdms[k] = CacheAdm{
			Id:                adm.Id,
			Aaid:              adm.Aaid,
			Title:             adm.Title,
			Status:            adm.Status,
			Type:              adm.Type,
			Groups:            adm.Groups,
			Position:          adm.Position,
			Purpose:           adm.Purpose,
			Material:          adm.Material,
			MaterialType:      adm.MaterialType,
			JumpUrl:           adm.JumpUrl,
			Region:            FormatRegion(adm.Region),
			Ranking:           adm.Ranking,
			Channel:           adm.Channel,
			Avatar:            adm.Avatar,
			Nickname:          adm.Nickname,
			Content:           adm.Content,
			DetailsUrl:        adm.DetailsUrl,
			ShowMembers:       adm.ShowMembers,
			CloseTime:         adm.CloseTime,
			CanClose:          adm.CanClose,
			Platform:          adm.Platform,
			PutType:           adm.PutType,
			StartTime:         adm.StartTime,
			EndTime:           adm.EndTime,
			Strategy:          formatStrategy(adm.Strategy),
			AthId:             adm.AthId,
			SplashTime:        adm.SplashTime,
			SplashForceTime:   adm.SplashForceTime,
			FrequencyType:     adm.FrequencyType,
			FrequencyCount:    adm.FrequencyCount,
			FrequencyInterval: adm.FrequencyInterval,
			FrequencyDays:     adm.FrequencyDays,
			ClientVersion:     formatClientVersion(adm.ClientVersion),
			PutTableName:      adm.PutTableName,
			SellType:          adm.SellType,
			InspiringTime:     adm.InspiringTime,
			ClickNum:          adm.ClickNum,
			InspireWatchNum:   adm.InspireWatchNum,
			IntervalNum:       adm.IntervalNum,
			InsertRow:         adm.InsertRow,
		}
	}

	return cAdms
}

type CacheClientVersion struct {
	CheckType    int
	StartVersion *version.Version
	EndVersion   *version.Version
}

func formatClientVersion(_version string) CacheClientVersion {
	var v ClientVersion
	var cv CacheClientVersion
	err := json.Unmarshal([]byte(_version), &v)
	if err != nil {
		return cv
	}

	cv.CheckType = v.CheckType
	if len(v.StartVersion) > 0 {
		cv.StartVersion, _ = version.NewVersion(v.StartVersion)
	}
	if len(v.EndVersion) > 0 {
		cv.EndVersion, _ = version.NewVersion(v.EndVersion)
	}

	return cv
}

//ranking 排序
func SortByRankingDesc(adms []CacheAdm) {
	sort.Slice(adms, func(i, j int) bool {
		return adms[i].Ranking > adms[j].Ranking
	})
}

//行数 排序
func SortByInsertRowAsc(adms []CacheAdm) {
	sort.Slice(adms, func(i, j int) bool {
		return adms[i].InsertRow < adms[j].InsertRow
	})
}

//地区码转map
type RegionsMap struct {
	Code  string
	Child *[]RegionsMap
}

func findRegion(code string, regions []RegionsMap) (bool, *[]RegionsMap) {
	for _, r := range regions {
		if r.Code == code {
			return true, r.Child
		}
	}

	return false, &[]RegionsMap{}
}

func insertRegionMap(region []string, k int, node *[]RegionsMap) {
	if k > 3 {
		return
	}
	var exists, tmpNode = findRegion(region[k], *node)
	//如果code不存在 且code不是0 则插入此code
	if exists == false && utils.Str2Int(region[k], 0) != 0 {
		*node = append(*node, RegionsMap{
			Code:  region[k],
			Child: tmpNode,
		})
	}
	insertRegionMap(region, k+1, tmpNode)
}

// region 1_156_110000,1_156_710000
func FormatRegion(region string) *[]RegionsMap {
	var res = &[]RegionsMap{}
	regions := strings.Split(region, ",")

	for _, rcode := range regions {
		tmpRs := strings.Split(rcode, "_")
		if len(tmpRs) < 3 {
			continue
		}
		tmp := tmpRs[2]
		rs := append(tmpRs[:2], tmp[:2], tmp[2:])

		insertRegionMap(rs, 0, res)

	}

	return res
}

//格式化策略 尾号
func formatStrategy(s string) Strategy {
	var strategy Strategy

	err := json.Unmarshal([]byte(s), &strategy)
	if err != nil {
		return strategy
	}

	return strategy
}
