package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		num := os.Args[i]
		numComma := num
		sign := ""
		dot := ""
		if strings.HasPrefix(num, "+") {
			numComma = strings.Replace(num, "+", "", 1)
			sign = "+"
		}
		if strings.HasPrefix(num, "-") {
			numComma = strings.Replace(num, "-", "", 1)
			sign = "-"
		}
		if inde := strings.Index(num, "."); inde != -1 {
			dot = numComma[inde:]
			numComma = numComma[:inde]
		}
		fmt.Printf("  %v%s%v\n", sign, comma(numComma), dot)
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