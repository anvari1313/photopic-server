package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	BasePath  string `mapstructure:"base_path"`
	URLPrefix string `mapstructure:"url_prefix"`
	Address   string `mapstructure:"address"`
}

var C *Config

func InitConfig() (*Config, error) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var c Config
	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	C = &c

	return &c, nil
}
