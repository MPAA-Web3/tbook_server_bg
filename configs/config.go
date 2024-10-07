package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var (
	config  GlobalConfig
	rConfig sync.RWMutex
)

// MysqlConfig mysql配置参数
type MysqlConfig struct {
	User     string
	Password string
	Ip       string
	Port     string
	DbName   string
}

// GlobalConfig 全局配置
type GlobalConfig struct {
	Port  string
	Mysql MysqlConfig
	Redis RedisConfig
	Debug bool
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

// Config 返回配置文件
func Config() GlobalConfig {
	rConfig.RLock()
	configCopy := config
	rConfig.RUnlock()
	return configCopy
}

func ParseConfig(cfg string) {
	viper.SetConfigFile(cfg)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("配置文件读取错误")
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
}
