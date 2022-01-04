package models

import (
	redisclient "git.ourbluecity.com/overseas-server-tools/go-util-redis"
	"github.com/ipipdotnet/ipdb-go"
	"gorm.io/gorm"
)

var (
	BluedDB          *gorm.DB
	RedisCli         redisclient.RedisClienter
	IPV4DB           *ipdb.City
	IPV6DB           *ipdb.City
	ArgoDomain       string
	ArgoBackupDomain string
	BackupDomain     string
	//UserImagesScheme       string
	ImagesDomain       string
	ImagesBackupDomain string
)

const (
	SCREEN_PROPORTION_SCALE          = 16.00
	INTER_3RD_AD_IMAGE_DEFAULT_VALUE = "http://7vznnm.com2.z0.glb.qiniucdn.com/blued-logo.png-500"
)

type ModelBase struct {
}
