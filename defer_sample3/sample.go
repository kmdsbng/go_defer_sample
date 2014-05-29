package main

import (
  "log"
)

func main() {
  started := make(chan struct{})
  stop := make(chan struct{})
  done := make(chan struct{})

  go func() {
    started <-struct{}{} // tell main thread we started
    log.Println("Goroutine starts")

    defer func() {
      log.Println("Defered func called!")
      done<-struct{}{} // tell main thread we're done
    }()

    <-stop // wait for main thread to tell us we should exit
  }()


  <-started // wait for the goroutine to start
  log.Println("Main thread receive Goroutine starts")
  stop<-struct{}{} // tell goroutine to stop
  <-done // wait for goroutine to be really done

  log.Println("Main thread exits")
}


