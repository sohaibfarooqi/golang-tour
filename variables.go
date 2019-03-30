package main

import "fmt"

func main(){
  var a int // Variable declaration var <variable-name> <dtype>
  fmt.Println(a)

  var b, c = 1, 2 // Variable declaration with initilizer
  fmt.Println(b,c)

  d := 3 // Short variable declaration
  fmt.Println(d)

  /*
  Zero value:
    - int --> 0
    - bool --> false
    - string --> ""
  */
  var e int
  var f bool
  var g string
  fmt.Println(e,f,g)

  h := 42
  i := float64(h) // Type conversion
  fmt.Println(i)

  const j = "Go lang" // Constant
  fmt.Println(j)

}
