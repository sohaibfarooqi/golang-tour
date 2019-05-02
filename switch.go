package main

import (
	"fmt"
	"runtime"
)

func main() {
	/*
	  Go switch statments run only the selected case and not all the cases
	  that follow i.e. Go inserts break statement itself. Cases are evaluated
	  from top to bottom.
	*/
	switch os := runtime.GOOS; os {

	case "darwin":
		fmt.Println("OS X")

	case "linux":
		fmt.Println("Linux.")

	default:
		fmt.Println(os)

	}
}
