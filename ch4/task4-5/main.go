package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 4, 4, 5, 6, 6, 6, 7, 7, 8}
	fmt.Println(remove(s))
}

func remove(slice []int) []int {
	minus := 0
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] == slice[i+1] {
			copy(slice[i:], slice[i+1:])
			minus++
			if i == len(slice)-minus {
				break
			}
			i--
		}
	}
	return slice[:len(slice)-minus+1]
}
