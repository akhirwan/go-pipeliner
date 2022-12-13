package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	if err := godotenv.Load(filenames...); err != nil {
		log.Printf("error: %v", err.Error())
		panic(err)
	}

	return &configImpl{}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
