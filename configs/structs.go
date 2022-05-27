package configs

type Config struct {
	Module  string `yaml:"Module"`
	Count   uint32 `yaml:"Count"`
	Input   string `yaml:"Input"`
	Threads uint16 `yaml:"Threads"`
	Secret  string `yaml:"Secret"`
}
