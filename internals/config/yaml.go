package config

import (
	"os"

	"github.com/githiago-f/lb-mini/internals/app"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Loadbalancer struct {
		Listen    int16
		Algorithm string
		Servers   []struct {
			Name string
			Host string
		}
	}
}

func Load() *Config {
	configData, err := os.ReadFile("./gobalancer.yml")
	if err != nil {
		app.Logger.Fatal(err)
		os.Exit(1)
	}

	res := &Config{}

	yaml.Unmarshal(configData, res)

	return res
}
