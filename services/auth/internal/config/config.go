package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env  string `yaml:"env"`
	Port int    `yaml:"port"`
}

var cfg *Config = nil

func MustRead(configPath string) *Config {
	op := "config.MustRead"

	c := Config{}
	err := cleanenv.ReadConfig(configPath, &c)

	if err != nil {
		panic(fmt.Sprintf("%s config file read error: %s", op, err.Error()))
	}
	cfg = &c

	return cfg
}

func Get() *Config {
	return cfg
}
