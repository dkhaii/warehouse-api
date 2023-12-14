package config

import (
	"github.com/spf13/viper"
)

type Config interface {
	GetString(key string) string
	GetInt(key string) int
}

type configImpl struct {
	viper *viper.Viper
}

func Init() (Config, error) {
	config := viper.New()

	config.SetConfigFile(".env")
	config.AddConfigPath("../")

	err := config.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &configImpl{
		viper: config,
	}, nil
}

func (config *configImpl) GetString(key string) string {
	return config.viper.GetString(key)
}

func (config *configImpl) GetInt(key string) int {
	return config.viper.GetInt(key)
}
