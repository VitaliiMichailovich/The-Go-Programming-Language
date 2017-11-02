package mypopcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func MyPopCount(x uint64) int {
	var ret byte
	var i uint
	for i = 0; i < 8; i++ {
		ret += pc[byte(x>>(i*8))]
	}
	return int(ret)
}