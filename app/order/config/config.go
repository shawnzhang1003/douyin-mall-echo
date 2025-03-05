package config

import (
	"github.com/MakiJOJO/douyin-mall-echo/common/mtl"
	"github.com/MakiJOJO/douyin-mall-echo/common/utils/config"
)

type Config struct {
	ENV      string                `mapstructure:"env"`
	HOST     string                `mapstructure:"host"`
	Log      config.Log            `mapstructure:"log"`
	Registry config.Registry       `mapstructure:"registry"`
	Kitex    config.Kitex          `mapstructure:"kitex"`
	JWT      config.JWTConfig      `mapstructure:"jwt"`
	MySQL    config.MySQL          `mapstructure:"mysql"`
	RocketMq config.ConsumerConfig `mapstructure:"rocketmq"`
}

var GlobalConfig Config

func Init(path string) error {
	// 初始化配置
	err := config.Init(path, &GlobalConfig)
	if err != nil {
		// 记录错误日志，但不直接终止程序
		mtl.Logger.Error("初始化配置失败", "configInitErr", err)
		return err
	}

	// 初始化成功，返回 nil
	return nil
}
