package conf

import "time"

type Config struct {
	Address        string        `yaml:"address"`
	Protocol       string        `yaml:"protocol"`
	RequestTimeout time.Duration `yaml:"request_timeout"`
}

func NewDefaultConfig() *Config {
	return &Config{
		RequestTimeout: 5 * time.Second,
	}
}

var Configs map[string]*Config

func init() {
	Configs = make(map[string]*Config)
}
