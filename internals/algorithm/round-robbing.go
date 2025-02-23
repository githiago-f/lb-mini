package algorithm

import "sync"

var (
	mu        = &sync.Mutex{}
	lastIndex = 0
)

func RoundRobbing(size int) uint32 {
	mu.Lock()
	res := ((lastIndex + 1) % size)
	lastIndex++
	mu.Unlock()
	return uint32(res)
}
