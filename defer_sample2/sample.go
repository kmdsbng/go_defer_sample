package main

import (
  "log"
)

func main() {
  started := make(chan struct{})
  stop := make(chan struct{})

  go func() {
    started <-struct{}{} // tell main thread we started
    log.Println("Goroutine starts")
    defer func() {
      log.Println("Defered func called!")
    }()

    <-stop // wait for main thread to tell us we should exit
  }()


  <-started // wait for the goroutine to start
  log.Println("Main thread exits")
}


