package grc

import (
	"fmt"

	"github.com/spf13/viper"
)

func Hello() {
	fmt.Println("Hello from GRC")
}

func LoadConfig(configName, configType, path string, conf *any) error {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&conf); err != nil {
		return err
	}
	return nil
}
