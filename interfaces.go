package main

import (
  "fmt"
  "math"
)

type Abser interface{
  Abs() float64
}

/*
`MyFloat` type implements `Abs` method from Abser interface
*/
type MyFloat float64
func (f MyFloat) Abs() float64{
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}

/*
`*Vertex` type implements `Abs` method from Abser interface
*/
type Vertex struct{
  X, Y float64
}
func (v *Vertex) Abs() float64{
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main(){
  var a Abser

  v := Vertex{3, 4}
  a = &v
  fmt.Println(a.Abs())

  f := MyFloat(-math.Sqrt2)
  a = f
  fmt.Println(a.Abs())
}
