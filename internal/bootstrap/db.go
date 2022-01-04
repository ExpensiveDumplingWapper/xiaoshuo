package bootstrap

import (
	"log"

	"xiaoshuo/internal/config"
	"xiaoshuo/internal/models"
	"xiaoshuo/pkg/db"
)

func loadDb() {
	conf, err := config.ZConf.GetMysqlConf("/blued/backend/udb/adm")
	models.BluedDB, err = db.NewMysqlConnection(db.Option{
		Master:   conf.Master,
		Slave:    conf.Slave,
		Username: conf.Username,
		Password: conf.Password,
		Database: "blued",
		//Port: 3306,
		Charset: "utf8",
	})

	if err != nil {
		log.Fatal(err.Error())
	}
}
