package main

import (
	"bufio"
	"fmt"
	"os"
)

type fill struct {
	file  []string
	count int
}

func main() {
	counts := make(map[string]fill)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.count > 1 {
			for _, v := range n.file {
				fmt.Printf("%d\t%s\t%v\n", n.count, line, v)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]fill) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] = fill{count: counts[input.Text()].count + 1, file: append(counts[input.Text()].file, string(f.Name()))}
	}
}
