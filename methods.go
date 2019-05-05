package main

import (
  "fmt"
  "math"
)

type Vertex struct{
  X, Y float64
}

/*
Go doesnot have classes, but still methods can
be defined on types.

A method is a func with special receiver argument
*/
func(v Vertex) Abs() float64{
  return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main(){
  v := Vertex{2,3}
  fmt.Println(v.Abs())
}
