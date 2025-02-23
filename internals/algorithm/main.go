package algorithm

import (
	"net/http/httputil"

	"github.com/githiago-f/lb-mini/internals/config"
)

func GetServer(remote string, config *config.Config) *httputil.ReverseProxy {
	size := len(config.Servers)

	var idx uint32
	switch config.Algorithm {
	case "RoundRobbing":
		idx = RoundRobbing(size)
		break
	case "HASH":
		idx = Hash(remote, size)
		break
	case "Weighted":
		break
	}

	return config.Servers[idx].Proxy
}
