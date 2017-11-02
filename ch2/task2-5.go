package main

import (
	"fmt"
	"github.com/VitaliiMichailovich/The-Go-Programming-Language/ch2/2-5-popcount"
	"math"
)

func main() {
	var a uint64 = math.MaxUint64
	fmt.Println(popcount.PopCount(a))
}
