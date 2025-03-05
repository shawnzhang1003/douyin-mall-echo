package config

import "time"

type ConsumerConfig struct {
	GroupName    string        `mapstructure:"groupname"`
	Topic        string        `mapstructure:"topic"`
	NsAddrs      []string      `mapstructure:"nsaddrs"`
	BrAddrs      []string      `mapstructure:"braddrs"`
	RetryCount   int           `mapstructure:"retrycount"`
	RetryTimeout time.Duration `mapstructure:"retrytimeout"`
}
