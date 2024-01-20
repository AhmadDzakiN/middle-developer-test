package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var defaultConfig *viper.Viper

func init() {
	defaultConfig = readViperConfig()
}

func readViperConfig() *viper.Viper {
	v := viper.New()
	v.AddConfigPath(".")
	v.AddConfigPath("../../../../params")
	v.AddConfigPath("./params")
	v.SetConfigName(".env")
	v.SetConfigType("env")

	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err == nil {
		fmt.Printf("Using config file: %s \n", v.ConfigFileUsed())
	} else {
		panic(fmt.Errorf("Config error: %s", err.Error()))
	}

	return v
}

func AppConfig() *viper.Viper {
	return defaultConfig
}
