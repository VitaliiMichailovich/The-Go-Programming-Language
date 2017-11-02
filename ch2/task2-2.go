package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/VitaliiMichailovich/The-Go-Programming-Language/ch2/2-2-conv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Need param to be converted.")
		os.Exit(1)
	}
	input, err := strconv.ParseFloat(os.Args[1], 100)
	if err != nil {
		fmt.Println(err.Error())
	}
	cels := conv.Celsius(input)
	fare := conv.Fahrenheit(input)
	kelv := conv.Kelvin(input)
	fmt.Println(cels, " is \t", conv.CToF(cels), "\tand\t", conv.CToK(cels))
	fmt.Println(fare, " is \t", conv.FToC(fare), "\tand\t", conv.FToK(fare))
	fmt.Println(kelv, " is \t", conv.KToF(kelv), "\tand\t", conv.KToC(kelv))
	foot := conv.Foot(input)
	metr := conv.Metr(input)
	fmt.Println(foot, " is \t", conv.FToM(foot))
	fmt.Println(metr, " is \t", conv.MToF(metr))
	pond := conv.Pond(input)
	kilo := conv.Kilo(input)
	fmt.Println(pond, " is \t", conv.PToK(pond))
	fmt.Println(kilo, " is \t", conv.KToP(kilo))
}