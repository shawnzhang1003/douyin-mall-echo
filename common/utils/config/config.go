package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init(path string, rawVal any) error {
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	// 实时监听config文件的变化
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(rawVal); err != nil {
			// handle error
			log.Printf("Error unmarshalling config: %v\n", err)
		}
	})
	viper.WatchConfig()

	if err := viper.Unmarshal(rawVal); err != nil {
		return err
	}

	return nil
}
