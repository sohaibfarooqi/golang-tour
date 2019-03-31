package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main(){
  fmt.Println(Vertex{1,2}) // Create a new Vertex with X=1, Y=2

  v := Vertex{1,2}
  fmt.Println(v.X) // Struct attribues are accessable using dot notation.

  p := &v // Reference struct v through pointer p
  p.X = 1e9 // Technically this should be (*p).X but Go allows easy syntax.
  fmt.Println(v)
}
