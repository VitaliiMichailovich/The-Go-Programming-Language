package main

import (
	"bufio"
	"fmt"
	"os"
)

type wordCount map[string]int

// run go run ch4/task4-9/main.go < ch4/task4-8/1.txt
func main() {
	counts := make(map[string]int) // counts of Unicode characters
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		counts[in.Text()]++
	}

	for k, v := range counts {
		fmt.Printf("%v : %v\n", k, v)
	}
}
