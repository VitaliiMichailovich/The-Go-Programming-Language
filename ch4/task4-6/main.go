package main

import (
	"fmt"
	"unicode"
)

func main() {
	in := []byte("Hi  eve   ry      one and to  you")
	fmt.Println(string(cutter(in)))
}

func cutter(input []byte) []byte {
	minus := 0
	for i := 0; i < len(input)-1; i++ {
		if unicode.IsSpace(rune(input[i])) && unicode.IsSpace(rune(input[i+1])) {
			copy(input[i:], input[i+1:])
			minus++
			if i == len(input)-minus+1 {
				break
			}
			i--
		}
	}
	return input[:len(input)-minus]
}
