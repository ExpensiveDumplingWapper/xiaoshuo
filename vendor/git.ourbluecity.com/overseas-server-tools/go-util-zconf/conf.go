package zconf

import (
	"github.com/go-zookeeper/zk"
)

func getConfPath(path string, conn *zk.Conn) (string, error) {
	confByte, _, err := conn.Get(path)
	if err != nil {
		return "", err
	}
	return string(confByte), nil
}

func getConfChildren(path string, conn *zk.Conn) ([]string, error) {
	host, _, err := conn.Children(path)
	if err != nil {
		return nil, err
	}
	returnData := make([]string, 0, 15)
	for _, v := range host {
		childByte, _, err := conn.Get(path + "/" + v)
		if err != nil {
			return nil, err
		}
		child := string(childByte)
		if child == "0" {
			returnData = append(returnData, v)
		}
	}
	return returnData, nil
}
