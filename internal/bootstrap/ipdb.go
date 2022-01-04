package bootstrap

import (
	"log"
	"os"

	"xiaoshuo/internal/config"
	"xiaoshuo/internal/models"

	"github.com/ipipdotnet/ipdb-go"
)

func loadIPDb() {
	var err error
	ipdbPath := os.Getenv("IPDB_PATH")
	var IpipData, IpipDataV6 = "", ""
	if ipdbPath != "" {
		IpipData = ipdbPath + "/city.ipv4.ipdb"
		IpipDataV6 = ipdbPath + "/city.ipv6.ipdb"
	} else {
		IpipData = config.Setting.Server.IpdbPath + "/city.ipv4.ipdb"
		IpipDataV6 = config.Setting.Server.IpdbPath + "/city.ipv6.ipdb"
	}
	models.IPV4DB, err = ipdb.NewCity(IpipData)
	if err != nil {
		log.Fatal(err)
	}

	models.IPV6DB, err = ipdb.NewCity(IpipDataV6)
	if err != nil {
		log.Fatal(err)
	}
}
