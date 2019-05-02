package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	/*
	  Creating a map
	*/
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.5, 89.2,
	}
	fmt.Println(m)

	/*
	  Map Literals
	*/
	var m1 = map[string]Vertex{
		"Bell Labs": Vertex{
			40.3, 89.3,
		},
		"Google": Vertex{
			65.4, 98.4,
		},
	}
	fmt.Println(m1)

	/*
	  Insert/Update key in map
	*/
	m1["Google"] = Vertex{10, 12}
	fmt.Println(m1)

	/*
	  Retrieve an element from map
	*/

	fmt.Println(m1["Google"])

	/*
	  Delete a key value pair
	*/

	delete(m1, "Google")
	fmt.Println(m1)

	/*
	  Test if a key is present
	*/

	elem, ok := m1["Google"]
	fmt.Println(elem, ok)

}
