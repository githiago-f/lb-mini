package algorithm

import "github.com/twmb/murmur3"

func Hash(remote string, size int) uint32 {
	return murmur3.Sum32([]byte(remote)) % uint32(size)
}
