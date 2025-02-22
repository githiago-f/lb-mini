package algorithm

var lastIndex = 0

func RoundRobbing(size int) uint32 {
	res := ((lastIndex + 1) % size)
	lastIndex++
	return uint32(res)
}
