package main
// defer_sample5 : close()を使うようにした
// defer_sample6 : sync.WaitGroupを使うようにした

import (
  "log"
  "sync"
)

func main() {
  const maxworkers int = 10

  started := &sync.WaitGroup {}
  stop := make(chan struct{})
  done := &sync.WaitGroup {}

  for i := 0; i < maxworkers; i++ {
    id := i // localize to be used in closure
    started.Add(1)
    done.Add(1)

    go func() {
      log.Printf("Goroutine %d starts", id)
      //started <-struct{}{} // tell main thread we started
      started.Done()

      defer func() {
        log.Printf("Defered func for %d called!", id)
        //done<-struct{}{} // tell main thread we're done
        done.Done()
      }()

      //time.Sleep(time.Duration(10 - id) * time.Second)
      <-stop // wait for main thread to tell us we should exit
    }()
  }

  started.Wait()
  log.Println("Main thread detect Groutine all runs")

  // do anything

  close(stop)

  done.Wait() // wait for goroutine to be really done

  log.Println("Main thread exits")
}


