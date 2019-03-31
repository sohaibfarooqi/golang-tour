package main

import "fmt"

func main(){
  /*
  defer function gets executed only when surrounding
  function returns.
  */
  defer fmt.Println("World")
  fmt.Println("Hello")

  /*
  Deferred functions are pushed onto stack. When function returns
  all pushed functions are invoked in last-in-first-out order.
  */
  for i := 0; i < 10; i++ {
    defer fmt.Println(i)
  }
}
