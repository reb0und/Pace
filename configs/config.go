package configs

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

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
