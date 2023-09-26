package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configImpl struct{}

func New(filenames ...string) (Config, error) {
	err := godotenv.Load(filenames...)
	if err != nil {
		return nil, err
	}

	return &configImpl{}, nil
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}
