package main

import (
	"encoding/json"
	"github.com/jinzhu/configor"
	"ipfc/eth"
	"ipfc/reward"
	"os"
)

const EnvConfigPath = "ENV_INSPECTOR_CONFIG_PATH"

var configFilePath = "conf/inspector.yaml"

func init() {
	path := os.Getenv(EnvConfigPath)
	if path != "" {
		configFilePath = path
	}
	loadConf()
}

type (
	IpfsConf struct {
		PeerId  string `yaml:"peer_id"`
		ApiAddr string `yaml:"api_addr"`
	}
	RepoConf struct {
		Dir string `yaml:"dir"`
	}
	MysqlConf struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password" json:"-"`
		Db       string `yaml:"db"`
	}
	Config struct {
		Mysql  MysqlConf     `yaml:"mysql"`
		Ipfs   IpfsConf      `yaml:"ipfs"`
		Repo   RepoConf      `yaml:"repo"`
		Eth    eth.Config    `yaml:"eth"`
		Reward reward.Config `yaml:"reward"`
	}
)

var AppConfig Config

func loadConf() {
	cfg := &Config{}
	err := configor.Load(cfg, configFilePath)
	if err != nil {
		panic(err)
		return
	}
	AppConfig = *cfg
}

func (c Config) String() string {
	data, _ := json.Marshal(c)
	return string(data)
}
