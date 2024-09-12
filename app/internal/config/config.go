package config

import (
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	IsDebug       bool `yaml:"is_debug"`
	IsDevelopment bool `yaml:"is_dev"`

	HTTP struct {
		IP           string        `yaml:"ip"`
		Port         string        `yaml:"port"`
		ReadTimeout  time.Duration `yaml:"read_timeout"`
		WriteTimeout time.Duration `yaml:"write_timeout"`
	} `yaml:"http"`

	AppConfig struct {
		LogLevel  string `yaml:"log_level"`
	} `yaml:"app_config"`

	PostgreSQL struct {
		Username string `yaml:"psql_username"`
		Password string `yaml:"psql_password"`
		Host     string `yaml:"psql_host"`
		Port     string `yaml:"psql_port"`
		Database string `yaml:"psql_database"`
	} `yaml:"postgresql"`
}

var instance *Config
var once sync.Once
var pathToConfig = "../../../configs/config.yaml"
var pathToConfigDocker = "config.yaml"

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig(pathToConfigDocker, instance); err != nil {
			helpText := "failed to read config"
			errText, _ := cleanenv.GetDescription(instance, &helpText)
			logrus.Error(errText)
			logrus.Fatal(err)
		}
	})
	return instance
}
