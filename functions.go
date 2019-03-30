package main

import "fmt"

/*
Function signature:
 - func <function-name>(<arg> <dtype>) <return type>{}
*/
func add(x int, y int) int{
  return x + y
}

/*
When two or more args of a function have same dtype.
We can ignore dtype of all args but last
*/
func sub(x, y int) int{
  return x - y
}

/*
A function can return any number of results
*/
func swap(x, y int) (int, int){
  return y, x
}

/*
Function return values can be named. They can be defined
at the top of function. At the end of function only `return`
statement is needed. This is called naked return statements.
*/
func split(sum int)(x, y int){
  x = sum  * 4/9
  y = sum - x
  return
}

func main(){
  fmt.Println(add(42, 13))
  fmt.Println(sub(42, 13))
  fmt.Println(swap(42, 13))
  fmt.Println(split(42))

}
