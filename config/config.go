package config

import (
	"project-go/exception"

	"github.com/spf13/viper"
)

type MainConfig struct {
	Database DatabaseConfig
	Server   ServerConfig
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
