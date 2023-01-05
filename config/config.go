package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	App     AppConfig     `yaml:"app"`
	HTTP    HTTPConfig    `yaml:"http"`
	Logger  LoggerConfig  `yaml:"logger"`
	Storage StorageConfig `yaml:"storage"`
}

type AppConfig struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type HTTPConfig struct {
	Port string `yaml:"port"`
}

type LoggerConfig struct {
	Level string `yaml:"log_level"`
}

type StorageConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var instance *Config
var err error

var once sync.Once

func GetConfig() (*Config, error) {
	once.Do(func() {
		instance = &Config{}
		err = cleanenv.ReadConfig("./config/config.yaml", instance)
		if err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
		}
	})
	if instance == nil {
		return nil, err
	}
	return instance, nil
}
