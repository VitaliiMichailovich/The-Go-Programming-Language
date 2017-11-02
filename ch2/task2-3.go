package main

import (
	"fmt"
	"github.com/VitaliiMichailovich/The-Go-Programming-Language/ch2/2-3-popcount"
	"github.com/VitaliiMichailovich/The-Go-Programming-Language/ch2/2-3-mypopcount"
)

func main() {
	var a uint64 = 7000
	fmt.Println(popcount.PopCount(a))
	fmt.Println(mypopcount.MyPopCount(a))
}
