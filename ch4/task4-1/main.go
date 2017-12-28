package main

import (
	"fmt"
	"crypto/sha256"
)

func main() {
	c1 := sha256.Sum256([]byte("hi"))
	c2 := sha256.Sum256([]byte("zzzzzzzz"))
	var i int
	for r := range c1 {
		if c1[r] == c2[r] {

			i++
		}
	}
	fmt.Println(i)
}