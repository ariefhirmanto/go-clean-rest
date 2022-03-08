package config

import (
	"project-go/exception"

	"github.com/spf13/viper"
)

type MainConfig struct {
	Database DatabaseConfig
	Server   ServerConfig
	Redis    CacheConfig
}

type ServerConfig struct {
	Address string `mapstructure:"Address"`
}

type DatabaseConfig struct {
	Host   string `mapstructure:"Host"`
	Port   string `mapstructure:"Port"`
	DBName string `mapstructure:"DBName"`
	DBUser string `mapstructure:"DBUser"`
	DBPass string `mapstructure:"DBPass"`
}

type CacheConfig struct {
	Host         string `mapstructure:"Host"`
	MinIdleConns int    `mapstructure:"MinIdleConns"`
	PoolSize     int    `mapstructure:"PoolSize"`
	PoolTimeout  int    `mapstructure:"PoolTimeout"`
	Password     string `mapstructure:"Password"`
	DB           int    `mapstructure:"DB"`
}

func LoadConfig(path string) (config MainConfig) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	exception.PanicIfNeeded(err)

	err = viper.Unmarshal(&config)
	exception.PanicIfNeeded(err)

	return config
}
