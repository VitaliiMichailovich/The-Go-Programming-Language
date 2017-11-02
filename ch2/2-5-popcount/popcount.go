package popcount

func PopCount(x uint64) int {
	i := 0
	for {
		if x == 0 { break }
		x = x & (x - 1)
		i++
	}
	return i
}
