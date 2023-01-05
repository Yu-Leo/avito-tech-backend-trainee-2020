package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	App     AppConfig     `yaml:"app"`
	Server  ServerConfig  `yaml:"server"`
	Logger  LoggerConfig  `yaml:"logger"`
	Storage StorageConfig `yaml:"storage"`
}

type AppConfig struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type ServerConfig struct {
	Host string `env-required:"true" yaml:"host"`
	Port int    `env-required:"true" yaml:"port"`
}

type LoggerConfig struct {
	Level string `yaml:"log_level"`
}

type StorageConfig struct {
	Host     string `env-required:"true" yaml:"host"`
	Port     int    `env-required:"true" yaml:"port"`
	Database string `env-required:"true" yaml:"database"`
	Username string `env-required:"true" yaml:"username"`
	Password string `env-required:"true" yaml:"password"`
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
