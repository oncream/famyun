package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	ServerConfig   ServerConfig   `mapstructure:"server"`
	LogConfig      LogConfig      `mapstructure:"log"`
	DatabaseConfig DatabaseConfig `mapstructure:"database"`
}

// ServerConfig Http服务器配置
type ServerConfig struct {
	Mode string `mapstructure:"mode"` // 可选 dev,prod, 默认 dev
	Port int    `mapstructure:"port"`
}

type LogConfig struct {
	Path  string `mapstructure:"path"`
	Level string `mapstructure:"level"`
}

type DatabaseConfig struct {
	Address  string            `mapstructure:"address"`
	Driver   string            `mapstructure:"driver"`
	Username string            `mapstructure:"username"`
	Password string            `mapstructure:"password"`
	Database string            `mapstructure:"database"`
	Params   map[string]string `mapstructure:"params"`
	ShowSql  bool              `mapstructure:"showsql"`
}

func NewAppConfig() *AppConfig {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file :%s", err)
	}
	viper.AutomaticEnv()

	config := &AppConfig{}
	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("unmarshal config err:%s", err)
	}
	return config
}
