package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// AppConfig config for application.
type AppConfig struct {
	Env  string `yaml:"env" env-default:"local"`
	GPRC GRPCConfig
}

// GRPCConfig config for GRPC server.
type GRPCConfig struct {
	Port int `yaml:"port"`
}

// DB config for Postgres Database
type DB struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	User   string `yaml:"user"`
	DBName string `yaml:"db_name"`
}

// MustLoad creates AppConfig and loads it from yaml file
func MustLoad() *AppConfig {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg AppConfig

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("config path is empty: " + err.Error())
	}

	return &cfg
}

// fetchConfigPath fetches config path from command line flag or environment variable.
// Priority: flag > env > default.
// Default value is empty string.
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
