package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Env           string `mapstructure:"env"`
	ServerAddress string `mapstructure:"server_address"`
	Debug         bool   `mapstructure:"debug"`
	GinMode       string `mapstructure:"gin_mode"`
	ServerUrl     string `mapstructure:"server_url"`
	DB_Driver     string `mapstructure:"db_driver"`
	DB_Hostname   string `mapstructure:"db_host"`
	DB_Port       int    `mapstructure:"db_port"`
	DB_Username   string `mapstructure:"db_user"`
	DB_Password   string `mapstructure:"db_password"`
	DB_Name       string `mapstructure:"db_name"`
}

var ServerUrl string

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigName("development")
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	ServerUrl = viper.GetString("server_url")

	err = viper.Unmarshal(&config)
	return
}
