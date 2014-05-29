package main

import (
  "log"
  "os"
)

func main() {
  file, err := os.Open("foo")
  if err != nil {
    log.Fatalf("Failed to open file: %s", err)
  }
  defer file.Close()
}


