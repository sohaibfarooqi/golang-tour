package main

import "fmt"

func main() {
	/*
	  Slice can be created from an array with syntax a[low: high].
	  Default value for low is zero and high is length of array.
	*/
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var m []int = primes[1:4]
	fmt.Println(m)

	/*
	  Slices only refer to underlying array. Changing the slice
	  modifies the original array.
	*/
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)

	a := arr[0:2]
	b := arr[1:3]

	b[0] = 5
	fmt.Println(a, b)
	fmt.Println(arr)

	/*
	  Slice literals are like arrays without length. The code below builds
	  an array then create a slice that reference it.
	*/
	r := []bool{true, false, true}
	fmt.Println(r)

	/*
	  A slice has both length and capacity. Lenght of slice is the number of
	  elements it contains and capacity of slice is number of elements in
	  underlying array starting from first element in slice.
	*/
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	s = s[:0]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	s = s[:2]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	s = s[2:]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	/*
	  Slice can be created using make(<dtype>, len, cap)
	*/
	x := make([]int, 0, 5)
	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)

	/*
	  Slice can be extended using append method.
	*/
	var y []int
	y = append(y, 1, 2, 3)
	fmt.Printf("len=%d cap=%d %v\n", len(y), cap(y), y)

}
