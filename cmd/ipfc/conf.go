package main

import (
	"encoding/json"
	"github.com/jinzhu/configor"
	"os"
)

const EnvConfigPath = "ENV_IPFC_CONFIG_PATH"

var configFilePath = "conf/ipfc.yaml"

func init() {
	path := os.Getenv(EnvConfigPath)
	if path != "" {
		configFilePath = path
	}
	loadConf()
}

type (
	LotusConf struct {
		ApiAddr string `yaml:"api_addr"`
		Token   string `yaml:"token"`
	}
	IpfsConf struct {
		ApiAddr string `yaml:"api_addr"`
	}
	HttpServerConf struct {
		ListenAddress string `yaml:"listen_address"`
	}
	Config struct {
		Lotus LotusConf      `yaml:"lotus"`
		Ipfs  IpfsConf       `yaml:"ipfs"`
		Http  HttpServerConf `yaml:"http"`
	}
)

var appConfig *Config

func loadConf() {
	cfg := &Config{}
	err := configor.Load(cfg, configFilePath)
	if err != nil {
		panic(err)
		return
	}
	appConfig = cfg
}

func (c *Config) String() string {
	data, _ := json.Marshal(c)
	return string(data)
}