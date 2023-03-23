package grc

import (
	"github.com/spf13/viper"
)

func LoadConfig[T any](configName, configType, path string, responseType T) (T, error) {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return responseType, err
	}

	var res T
	if err := viper.Unmarshal(&res); err != nil {
		return responseType, err
	}
	return res, nil
}
