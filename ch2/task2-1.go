package main

import (
	"fmt"
	"github.com/VitaliiMichailovich/The-Go-Programming-Language/ch2/2-1-tempconv"
)

func main() {
	var cels tempconv.Celsius = 100
	var ctof tempconv.Fahrenheit = tempconv.CToF(cels)
	var ftok tempconv.Kelvin = tempconv.FToK(ctof)
	var ktoc tempconv.Celsius = tempconv.KToC(ftok)
	fmt.Println(cels, ctof, ftok, ktoc)
}