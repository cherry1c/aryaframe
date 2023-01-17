package conf

import "time"

type Config struct {
	Grpc         GrpcConfig    `yaml:"grpc"`
	Http         HttpConfig    `yaml:"http"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

type GrpcConfig struct {
	Address string `yaml:"address"`
}

type HttpConfig struct {
	Address string `yaml:"address"`
}

func NewDefaultConfig() *Config {
	return &Config{
		Grpc:         GrpcConfig{},
		Http:         HttpConfig{},
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
}
