package main

import "fmt"

func main(){
  /*
  Go has only one loop i.e. `for` loop.
  */
  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)

  /*
  Init and post condition are optional. That means you can
  create a while loop like below
  */
  sum_ := 1
  for sum_ < 1000 {
    sum_ += sum_
  }
  fmt.Println(sum_)
}
