package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Logger     LoggerConf     `yaml:"logger"`
	Database   DBConf         `yaml:"database"`
	HTTPServer HTTPServerConf `yaml:"httpServer"`
}

type LoggerConf struct {
	Level string `yaml:"level"`
}

type DBConf struct {
	Storage       string `yaml:"storage"`
	ConnectString string `yaml:"connectStr"`
}

type HTTPServerConf struct {
	Address string `yaml:"address"`
}

func NewConfig(configFile string) Config {
	config := Config{}
	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal("error reading the configuration file")
	}

	yaml.Unmarshal(yamlFile, &config)
	return config
}
