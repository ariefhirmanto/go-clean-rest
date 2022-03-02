package config

import (
	"fmt"

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

func LoadConfig(path string) (config MainConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Failed to unmarshal")
		return config, err
	}

	fmt.Printf("%+v\n", config)
	return config, nil
}
