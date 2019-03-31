package main

import "fmt"

func main(){

  a, b := 1,2

  c := &a // c points to a
  fmt.Println(*c)

  *c = 3 // set a through c
  fmt.Println(a)

  c = &b // c now points to b
  *c = *c / 2 // set values of b through c
  fmt.Println(b)

}
