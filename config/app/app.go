package app

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var Conf *ini.File

func configRead() *ini.File {
	cfg, err := ini.Load("env.ini")
	if err != nil {
		fmt.Printf("fail to read ini file:%v\n", err)
	}
	return cfg
}

func init() {
	Conf = configRead()
}

func Config(section string, key string) string {
	return Conf.Section(section).Key(key).String()
}
