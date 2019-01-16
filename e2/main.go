package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

const (
	kbSize = 1024
)

func main() {
	var dnsAddr string
	flag.StringVar(&dnsAddr, "addr", "8.8.8.8", "the address of the DNS server")
	flag.Parse()

	addr := fmt.Sprintf("%s:80", dnsAddr)
	fmt.Printf("dailing %s...\n\n", addr)
	conn, err := net.Dial("tcp", addr)
	handleErr(err)

	defer func() {
		closeErr := conn.Close()
		handleErr(closeErr)
	}()

	req := []byte("HEAD / HTTP/1.1\r\n\r\n")
	m, err := conn.Write(req)
	handleErr(err)
	fmt.Printf("wrote %d bytes\n", m)
	fmt.Printf("%s\n", req)

	res := make([]byte, kbSize*kbSize)
	for {
		n, readErr := conn.Read(res)
		handleErr(readErr)
		fmt.Printf("read %d bytes\n", n)
		fmt.Printf("%s\n", res)

		time.Sleep(time.Second)
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
