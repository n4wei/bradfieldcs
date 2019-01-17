package main

import (
  "fmt"
  "net/http"
  "strings"
  "log"
)

func main() {
  http.HandleFunc("/", forwardReq)
  err := http.ListenAndServe(":3333", nil)
  if err != nil {
		panic(err)
  }
}

func forwardReq(w http.ResponseWriter, r *http.Request) {
}

