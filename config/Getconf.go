package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	MYSQL struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Addr     string `yaml:"addr"`
		Database string `yaml:"database"`
	}
	Server struct {
		Port string `yaml:"port"`
	}
	Send struct {
		Cookie     string `yaml:"cookie"`
		Referer    string `yaml:"referer"`
		XCsrftoken string `yaml:"x-csrftoken"`
	}
}

//获取配置
func GetConf() *Conf {
	var c = Conf{}
	yamlFile, err := ioutil.ReadFile("./config/conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return &c
}

//获取配置文件
var Config = GetConf()
