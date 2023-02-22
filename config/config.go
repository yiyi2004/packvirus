package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	ModelPath   string `yaml:"model_path"`
	PayloadPath string `yaml:"payload_path"`
	Key         string `yaml:"key"`
	Algorithm   string `yaml:"algorithm"`
}

func LoadConfig(path string) (*Config, error) {
	// 判断是否是 path
	// 用 yaml 加载配置文件并返回

	var config Config
	configFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
