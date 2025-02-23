package algorithm

import "sync"

var (
	mu        = &sync.Mutex{}
	lastIndex = 0
)

func RoundRobbing(size int) uint32 {
	mu.Lock()
	defer mu.Unlock()
	res := ((lastIndex + 1) % size)
	lastIndex++
	return uint32(res)
}
