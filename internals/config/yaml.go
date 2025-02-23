package config

import (
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/githiago-f/lb-mini/internals/app"
	"gopkg.in/yaml.v3"
)

type ConfigFileData struct {
	Loadbalancer struct {
		Listen    int16
		Algorithm string
		Servers   []struct {
			Host string
			Name string
		}
	}
}

type Server struct {
	Proxy *httputil.ReverseProxy
}
type Config struct {
	Listen    int16
	Algorithm string
	Servers   []*Server
}

func Load() *Config {
	configData, err := os.ReadFile("./gobalancer.yml")
	if err != nil {
		app.Logger.Fatal(err)
		os.Exit(1)
	}

	data := &ConfigFileData{}

	yaml.Unmarshal(configData, data)

	res := &Config{}
	res.Algorithm = data.Loadbalancer.Algorithm
	res.Listen = data.Loadbalancer.Listen
	res.Servers = make([]*Server, len(data.Loadbalancer.Servers))

	for i, server := range data.Loadbalancer.Servers {
		app.Logger.Infof("Included %v to servers list", server.Name)
		url, err := url.Parse(server.Host)
		if err != nil {
			app.Logger.Fatal(err)
			os.Exit(1)
		}
		res.Servers[i] = &Server{Proxy: httputil.NewSingleHostReverseProxy(url)}
	}

	return res
}
