package mysql

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type DBConfig struct {
	User     string
	Password string
	Proto    string
	HostIP   string
	Port     int
	DBName   string
}

func (dc DBConfig) String() string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s", dc.User, dc.Password, dc.Proto, dc.HostIP, dc.Port, dc.DBName)
}

func ReadConfig() (*DBConfig, error) {
	var x DBConfig
	data, err := ioutil.ReadFile("./mysql/config.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return nil, err
	}
	return &x, nil
}
