package zconf

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-zookeeper/zk"
)

const (
	connectTimeOut = 10
)

type cacheData struct {
	time         time.Time
	mysqlConf    MysqlConf
	childrenConf []string
	conf         string
}

// ZConf 配置选项
type ZConf struct {
	Path     []string
	Duration time.Duration
	pathData sync.Map
	once     sync.Once
}

func (z *ZConf) setPathData(path string, data cacheData) {
	z.pathData.Store(path, data)
}

func (z *ZConf) getPathData(path string, now time.Time) (cacheData, bool, bool) {
	oldData, ok := z.pathData.Load(path)
	if !ok {
		return cacheData{}, false, false
	}
	data, typeOk := oldData.(cacheData)
	if !typeOk {
		return cacheData{}, false, false
	}
	if now.Before(data.time.Add(z.Duration)) {
		return data, true, true
	}
	return data, false, true
}

// GetMysqlConf 获取mysql配置
func (z *ZConf) GetMysqlConf(path string) (MysqlConf, error) {
	now := time.Now()
	oldData, notTimout, ok := z.getPathData(path, now)
	if notTimout && ok {
		return oldData.mysqlConf, nil
	}
	conn, _, err := zk.Connect(z.Path, time.Second*connectTimeOut)
	if conn != nil {
		defer conn.Close()
	}
	if err != nil {
		errMessage := fmt.Errorf("zk[mysql] connect error, error is %v", err)
		if ok {
			fmt.Printf("%v\n", errMessage)
			return oldData.mysqlConf, nil
		}
		return MysqlConf{}, errMessage
	}
	mysqlConf, err := getMysqlPath(path, conn)
	if err != nil {
		errMessage := fmt.Errorf("zk[mysql] get path error, error is %v", err)
		if ok {
			fmt.Printf("%v\n", errMessage)
			return oldData.mysqlConf, nil
		}
		return MysqlConf{}, errMessage
	}
	z.setPathData(path, cacheData{
		time:      now,
		mysqlConf: mysqlConf,
	})
	return mysqlConf, nil
}

// GetConfChildren 获取子目录配置
func (z *ZConf) GetConfChildren(path string) ([]string, error) {
	now := time.Now()
	oldData, notTimout, ok := z.getPathData(path, now)
	if notTimout && ok {
		return oldData.childrenConf, nil
	}
	conn, _, err := zk.Connect(z.Path, time.Second*connectTimeOut)
	if conn != nil {
		defer conn.Close()
	}
	if err != nil {
		errMessage := fmt.Errorf("zk[childrenConf] connect error, error is %v", err)
		if ok {
			fmt.Printf("%v\n", errMessage)
			return oldData.childrenConf, nil
		}
		return nil, errMessage
	}
	childrenConf, err := getConfChildren(path, conn)
	if err != nil {
		errMessage := fmt.Errorf("zk[childrenConf] get path error, error is %v", err)
		if ok {
			fmt.Printf("%v\n", errMessage)
			return oldData.childrenConf, nil
		}
		return nil, errMessage
	}
	z.setPathData(path, cacheData{
		time:         now,
		childrenConf: childrenConf,
	})
	return childrenConf, nil
}

// GetConf 获取键值对配置
func (z *ZConf) GetConf(path string) (string, error) {
	now := time.Now()
	oldData, notTimout, ok := z.getPathData(path, now)
	if notTimout && ok {
		return oldData.conf, nil
	}
	conn, _, err := zk.Connect(z.Path, time.Second*connectTimeOut)
	if conn != nil {
		defer conn.Close()
	}
	if err != nil {
		errMessage := fmt.Errorf("zk[conf] connect error, error is %v", err)
		if ok {
			fmt.Printf("%v\n", errMessage)
			return oldData.conf, nil
		}
		return "", errMessage
	}
	conf, err := getConfPath(path, conn)
	if err != nil {
		errMessage := fmt.Errorf("zk[conf] get path error, error is %v", err)
		if ok {
			fmt.Printf("%v\n", errMessage)
			return oldData.conf, nil
		}
		return "", errMessage
	}
	z.setPathData(path, cacheData{
		time: now,
		conf: conf,
	})
	return conf, nil
}
