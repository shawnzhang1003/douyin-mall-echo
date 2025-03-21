package config

import (
	"log"

	"github.com/MakiJOJO/douyin-mall-echo/common/utils/config"
)

type Config struct {
	ENV      string           `mapstructure:"env"`
	HOST     string           `mapstructure:"host"`
	Log      config.Log       `mapstructure:"log"`
	Registry config.Registry  `mapstructure:"registry"`
	Kitex    config.Kitex     `mapstructure:"kitex"`
	JWT      config.JWTConfig `mapstructure:"jwt"`
	MySQL    config.MySQL     `mapstructure:"mysql"`
	Redis    config.Redis     `mapstructure:"redis"`
}

var GlobalConfig Config

func Init(path string) {
	if err := config.Init("../config/config.yaml", &GlobalConfig); err != nil {
		log.Fatalf("初始化配置失败: %v", err)
	}
}
