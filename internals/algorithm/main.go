package algorithm

import (
	"net/url"
	"os"

	"github.com/githiago-f/lb-mini/internals/app"
	"github.com/githiago-f/lb-mini/internals/config"
)

func GetServer(remote string, config *config.Config) *url.URL {
	size := len(config.Loadbalancer.Servers)

	var idx uint32
	switch config.Loadbalancer.Algorithm {
	case "RoundRobbing":
		idx = RoundRobbing(size)
		break
	case "HASH":
		idx = Hash(remote, size)
		break
	case "Weighted":
		break
	}

	host := config.Loadbalancer.Servers[idx].Host
	app.Logger.Debugf(":: %s", host)

	res, err := url.Parse(host)
	if err != nil {
		app.Logger.Fatalf("%v while trying to parse %s", err, host)
		os.Exit(1)
	}

	return res
}
