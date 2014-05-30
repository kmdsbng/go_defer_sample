package main
// defer_sample5 : close()を使うようにした
// defer_sample6 : sync.WaitGroupを使うようにした
// defer_sample8 : groutineの戻り値も取る

import (
  "log"
  "sync"
)

func main() {
  const maxworkers int = 10

  started := &sync.WaitGroup {}
  stop := make(chan struct{})
  done := &sync.WaitGroup {}
  results := make(chan int)

  for i := 0; i < maxworkers; i++ {
    id := i // localize to be used in closure
    started.Add(1)
    done.Add(1)

    go func() {
      started.Done()
      log.Printf("Goroutine %d starts", id)

      defer func() {
        log.Printf("Defered func for %d called!", id)
        done.Done()
      }()

      <-stop // wait for main thread to tell us we should exit

      log.Printf("Groutine %d received stop", id)
      results<-id // pass our result
    }()

  }

  started.Wait()
  log.Println("Main thread detect Groutine all runs")

  go func() {
    done.Wait()
    log.Printf("Closing results")
    close(results)
  }()

  close(stop)

  ids := []int {}
  for v := range results {
    ids = append(ids, v)
  }

  done.Wait() // wait for goroutine to be really done

  log.Printf("Got %#d", ids)
  log.Println("Main thread exits")
}


