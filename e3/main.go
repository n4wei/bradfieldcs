package main

import (
  "fmt"
  "net/http"
	"os"
	"flag"
	"io/ioutil"
)

var forwardedPort string

func main() {
	var listenPort string
	flag.StringVar(&listenPort, "l", "3000", "the server will listen on this port")
	flag.StringVar(&forwardedPort, "f", "3001", "the port on the localhost to forward to")
	flag.Parse()

  http.HandleFunc("/", forwardReq)
	fmt.Printf("listening on localhost:%s\n", listenPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s",listenPort), nil)
  if err != nil {
		handleServerErr(err)
  }
}

func forwardReq(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	forwardedReq, err := http.NewRequest(r.Method, fmt.Sprintf("http://localhost:%s", forwardedPort), r.Body)
	handleServerErr(err)
	res, err := client.Do(forwardedReq)
	handleServerErr(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	handleServerErr(err)
	n, err := w.Write(body)
	handleServerErr(err)
	fmt.Printf("wrote %d bytes\n", n)
}

func handleServerErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
