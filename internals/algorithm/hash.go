package algorithm

import (
	"strings"

	"github.com/githiago-f/lb-mini/internals/app"
	"github.com/twmb/murmur3"
)

func Hash(remote string, size int) uint32 {
	host := strings.Split(remote, ":")[0]
	app.Logger.Debugf("remote :: %s", host)

	hash := murmur3.Sum32([]byte(host))
	app.Logger.Debugf("hash :: %d", hash)

	return hash % uint32(size)
}
