package main

import (
  "fmt"
  "time"
  "strconv"
  "math/rand"
)

func f(n int){
  for i := 0; i < 10; i++{
    fmt.Println(n, ":", i)
    amt := time.Duration(rand.Intn(250))
    time.Sleep(time.Millisecond * amt)
  }
}

func Ping(c chan string){
  for i := 0; i < 10; i++{
    c <- "ping" + strconv.Itoa(i)
  }
}

func Pong(c chan string) {
  for i := 0; ; i++ {
    c <- "pong" + strconv.Itoa(i)
  }
}

func StdOut(c chan string){
  for {
    msg := <- c
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
}

func main(){
  go f(0) // Running as Goroutine

  /*
  Using channels
  */
  var c chan string = make(chan string)
  go Ping(c)
  go Pong(c)
  go StdOut(c)

  /*
  Using select statement
  */
  c1 := make(chan string)
  c2 := make(chan string)

  go func(){
    for{
      c1 <- "from 1"
      time.Sleep(time.Second * 2)
    }
  }()

  go func(){
    for{
      c2 <- "from 2"
      time.Sleep(time.Second * 3)
    }
  }()

  go func(){
    for{
      select{
      case msg1 := <- c1:
        fmt.Println(msg1)
      case msg2 := <- c2:
        fmt.Println(msg2)
      }
    }
  }()



  var input string
  fmt.Scanln(&input)
}
