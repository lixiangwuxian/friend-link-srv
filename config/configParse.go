package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var config *Config

func GetConfig() *Config {
	if config == nil {
		getConfig()
	}
	return config
}

func getConfig() (*Config, error) {
	// 读取YAML文件内容
	yamlFile, err := os.ReadFile("conf.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}
	// 解析YAML内容到Config结构体
	err = yaml.Unmarshal(yamlFile, &config)
	return config, err
}
