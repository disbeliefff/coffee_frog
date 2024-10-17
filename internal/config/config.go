package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ServerPort int      `yaml:"server_port"`
	DB         DBConfig `yaml:"db"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func Load(filePath string) *Config {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic("Failed to read config file: " + err.Error())
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic("Failed to parse config file: " + err.Error())
	}

	return &cfg
}
