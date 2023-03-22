package grc

import (
	"fmt"

	"github.com/spf13/viper"
)

func Hello() {
	fmt.Println("Hello from GRC")
}

func LoadConfig(path string, conf interface{}) (c interface{}, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(conf)
	c = conf
	return c, err
}
