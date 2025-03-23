package rocketmq

import (
    "github.com/apache/rocketmq-client-go/v2/primitive"
)

// Config 定义 RocketMQ 配置结构体
type Config struct {
    NameServerAddr []string
    GroupName      string
}

// NewConfig 创建一个新的 RocketMQ 配置实例
func NewConfig(nameServerAddr []string, groupName string) *Config {
    return &Config{
        NameServerAddr: nameServerAddr,
        GroupName:      groupName,
    }
}

// GetNsResolver 获取 RocketMQ 命名服务器解析器
func (c *Config) GetNsResolver() primitive.NsResolver {
    return primitive.NewPassthroughResolver(c.NameServerAddr)
}