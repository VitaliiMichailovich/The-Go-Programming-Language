package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) (retun string) {
	n := len(s)
	if n <= 3 {
		return s
	}
	if n%3 != 0 {
		retun = s[:n-(n/3)*3]+","
	}
	for i := n/3; i >= 1; i-- {
		pos := n-i*3
		retun += s[pos:pos+3]
		if i!=1 {
			retun += ","
		}
	}
	return retun
}