package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Cfg struct {
		Services Services `yaml:"services"`
		Port     string   `env:"PORT"`
	}

	Services struct {
		Auth    SvcConfig `yaml:"auth"`
		User    SvcConfig `yaml:"user"`
		Project SvcConfig `yaml:"project"`
		Task    SvcConfig `yaml:"task"`
	}

	SvcConfig struct {
		Host string `yaml:"host"`
	}

	option func(*Cfg)
)

func New() *Cfg {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "development" {
		return load(
			readFrom(
				"./config/development.env",
				"./config/development.services.yaml",
			),
		)
	}

	if appEnv == "production" {
		return load(
			readFrom(
				"./config/production.env",
				"./config/production.services.yaml",
			),
		)
	}

	return nil
}

func load(opts ...option) *Cfg {
	var cfg Cfg
	for _, option := range opts {
		option(&cfg)
	}

	return &cfg
}

func readFrom(paths ...string) option {
	return func(c *Cfg) {
		for _, path := range paths {
			err := cleanenv.ReadConfig(path, c)
			if err != nil {
				panic(err)
			}
		}
	}
}
