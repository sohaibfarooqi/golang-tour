package main

import "fmt"

func main() {
	// Array Syntax var <name> [<number of elements>]<dtype>
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])

	// assigning a constant array to primes.
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
