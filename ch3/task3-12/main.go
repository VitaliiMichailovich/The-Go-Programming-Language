package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("I need two arguments.")
		return
	}
	anagramma := os.Args[1]
	ammargana := os.Args[2]
	if len(anagramma) != len(ammargana) {
		fmt.Println(anagramma + " is not an anagramma of " + ammargana)
		return
	}
	for _, l := range anagramma {
		anagramma = strings.Replace(anagramma, string(l), "", -1)
		ammargana = strings.Replace(ammargana, string(l), "", -1)
	}
	if len(anagramma) == len(ammargana) && anagramma == ammargana {
		fmt.Println("This two words is an anagramma of each other")
	} else {
		fmt.Println("This two words is NOT an anagramma of each other")
	}
}
