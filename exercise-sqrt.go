package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	y:= 0.0
	i := 0
	for  z != y {
		y=z
		z -= (z*z - x) / (2*z)
		i++
		if i > 100 {return z}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(9))
}
