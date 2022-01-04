package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func NewMysqlConnection(option Option) (*gorm.DB, error) {
	dialectors := getDialectors(option)
	db, err := gorm.Open(dialectors.Master, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	config := dbresolver.Config{
		Replicas: dialectors.Slave,
		Policy:   dbresolver.RandomPolicy{},
	}

	err = db.Use(
		dbresolver.Register(config).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(10).
			SetMaxOpenConns(20),
	)

	if err != nil {
		return nil, err
	}
	return db, nil
}

type Option struct {
	Master   string
	Slave    []string
	Username string
	Password string
	//Port int64
	Database string
	Charset  string
}

type MasterSlaveDialector struct {
	Master gorm.Dialector
	Slave  []gorm.Dialector
}

func getDialectors(option Option) MasterSlaveDialector {

	fmt.Println(option)
	var slaveDialector []gorm.Dialector

	masterDsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		option.Username, option.Password, option.Master, option.Database, option.Charset)

	for _, v := range option.Slave {
		var slaveDsn string
		slaveDsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
			option.Username, option.Password, v, option.Database, option.Charset)

		slaveDialector = append(slaveDialector, mysql.Open(slaveDsn))
	}
	masterDialector := mysql.Open(masterDsn)

	return MasterSlaveDialector{Master: masterDialector, Slave: slaveDialector}
}
