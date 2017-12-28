package main

import (
	"flag"
	"fmt"
	"os"
	"crypto/sha256"
	"crypto/sha512"
)

func main() {
	shaInt := flag.Int("sha", 256, "an int")
	flag.Parse()
	sha := *shaInt
	if len(os.Args) == 1 {
		fmt.Println("Put the string to be hashed as an argument and rerun the programm.")
		return
	}
	str := os.Args[1]
	str2 := os.Args[2]
	if string(str[0]) == "-" && string(str2[0]) != "-" {
		str = os.Args[2]
	}
	switch sha {
	case 384:
		fmt.Printf("%x", sha512.Sum384([]byte(str)))
	case 512:
		fmt.Printf("%x", sha512.Sum512([]byte(str)))
	default:
		fmt.Printf("%x", sha256.Sum256([]byte(str)))
	}
}
