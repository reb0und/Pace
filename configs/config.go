package configs

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Module  string `yaml:"Module"`
	Count   uint32 `yaml:"Count"`
	Input   string `yaml:"Input"`
	Threads uint16 `yaml:"Threads"`
	Secret  string `yaml:"Secret"`
}

func LoadConfig(path string) (*Config, error) {
	bytes, err := os.ReadFile(fmt.Sprintf("bin/%s", path))
	if err != nil {
		return nil, err
	}

	data := Config{}

	if err := yaml.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
