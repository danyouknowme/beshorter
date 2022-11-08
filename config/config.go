package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Env           string `mapstructure:"env"`
	ServerAddress string `mapstructure:"server_address"`
	Debug         bool   `mapstructure:"debug"`
	GinMode       string `mapstructure:"gin_mode"`
	DB_Driver     string `mapstructure:"db_driver"`
	DB_Source     string `mapstructure:"db_source"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigName("development")
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
