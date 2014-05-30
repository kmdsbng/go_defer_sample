package main

import (
  "log"
)

func main() {
  const maxworkers int = 10

  started := make(chan struct{})
  stop := make(chan struct{})
  done := make(chan struct{})

  for i := 0; i < maxworkers; i++ {
    id := i // localize to be used in closure

    go func() {
      log.Printf("Goroutine %d starts", id)
      started <-struct{}{} // tell main thread we started

      defer func() {
        log.Printf("Defered func for %d called!", id)
        done<-struct{}{} // tell main thread we're done
      }()

      //time.Sleep(time.Duration(10 - id) * time.Second)
      <-stop // wait for main thread to tell us we should exit
    }()
  }

  for i := 0; i < maxworkers; i++ {
    <-started // wait for the goroutine to start
  }
  log.Println("Main thread detect Groutine all runs")

  // do anything

  close(stop)

  for i := 0; i < maxworkers; i++ {
    <-done // wait for goroutine to be really done
  }

  log.Println("Main thread exits")
}


