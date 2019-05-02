package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	/*
	  if statement can have a short statement before actual condition.
	  Variables declared in short statement are also available in else.
	*/
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Println("%g >= %g\n", v, lim)
	}
	return lim
}
func main() {
	fmt.Println(sqrt(2), sqrt(-2))
	fmt.Println(pow(3, 2, 10))
}
