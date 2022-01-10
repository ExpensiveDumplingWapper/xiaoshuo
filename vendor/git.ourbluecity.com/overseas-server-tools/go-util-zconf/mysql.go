package zconf

import (
	"github.com/go-zookeeper/zk"
)

// MysqlConf mysql 配置选项
type MysqlConf struct {
	Master   string
	Slave    []string
	Username string
	Password string
}

func getMysqlPath(path string, conn *zk.Conn) (MysqlConf, error) {
	var (
		userNamePath = path + "/username"
		passwordPath = path + "/password"
		masterPath   = path + "/master"
		slavePath    = path + "/slave"
	)

	userName, _, err := conn.Get(userNamePath)
	if err != nil {
		return MysqlConf{}, err
	}
	passwrod, _, err := conn.Get(passwordPath)
	if err != nil {
		return MysqlConf{}, err
	}
	master, _, err := conn.Children(masterPath)
	if err != nil {
		return MysqlConf{}, err
	}
	slave, _, err := conn.Children(slavePath)
	if err != nil {
		return MysqlConf{}, err
	}
	returnData := MysqlConf{
		Username: string(userName),
		Password: string(passwrod),
		Master:   master[0],
		Slave:    make([]string, 0, 10),
	}

	for _, v := range slave {
		childSlaveByte, _, err := conn.Get(slavePath + "/" + v)
		if err != nil {
			return MysqlConf{}, err
		}
		childSlave := string(childSlaveByte)
		if childSlave == "0" {
			returnData.Slave = append(returnData.Slave, v)
		}
	}
	for _, v := range master {
		childMasterByte, _, err := conn.Get(masterPath + "/" + v)
		if err != nil {
			return MysqlConf{}, err
		}
		childMaster := string(childMasterByte)
		if childMaster == "0" {
			returnData.Master = v
			break
		}
	}

	return returnData, nil
}
