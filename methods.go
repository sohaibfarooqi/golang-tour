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

/*
Methods can be declared on non-struct types too.
*/
type MyFloat float64

func (f MyFloat) Abs() float64{
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}

/*
Receiver argument can be a pointer too.
This will modify the original argumnet passed
instead of modifying a copy.
*/
func (v *Vertex) Scale(f float64){
  v.X = v.X * f
  v.Y = v.Y * f
}
func main(){
  v := Vertex{2,3}
  fmt.Println(v.Abs())

  f := MyFloat(-math.Sqrt2)
  fmt.Println(f.Abs())

  // For convenience `v.Scale(..)` is acceptable instead of (&v).Scale(..)
  v.Scale(10)
  fmt.Println(v.Abs())

}
