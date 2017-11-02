package main

import (
	"time"
	"fmt"
	"strings"
	"os"
)

func main() {
	now := time.Now()
	fmt.Println(strings.Join(os.Args[:], " "))
	fmt.Println(time.Since(now))
	now = time.Now()
	for k, v := range os.Args {
		fmt.Println(k, " : ", v)
	}
	fmt.Println(time.Since(now))
}
