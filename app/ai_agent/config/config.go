package config

import (
	"log"

	"github.com/MakiJOJO/douyin-mall-echo/common/utils/config"
)

type Config struct {
	ENV      string           `mapstructure:"env"`
	HOST     string           `mapstructure:"host"`
	Registry config.Registry  `mapstructure:"registry"`
	Kitex    config.Kitex     `mapstructure:"kitex"`
	JWT      config.JWTConfig `mapstructure:"jwt"`
	MySQL    config.MySQL     `mapstructure:"mysql"`
	ApiKey   string           `mapstructure:"api_key"`
	Model    string           `mapstructure:"model"`
}

var GlobalConfig Config

func Init() {
	if err := config.Init("../config/config.yaml", &GlobalConfig); err != nil {
		log.Fatalf("初始化配置失败: %v", err)
	}
}
