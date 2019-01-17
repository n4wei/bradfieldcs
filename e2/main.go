package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

const (
	kbSize = 1024
)

func main() {
	var dnsAddr string
	var resolveAddr string
	flag.StringVar(&dnsAddr, "dns", "8.8.8.8", "the address of the DNS server")
	flag.StringVar(&resolveAddr, "addr", "", "the address to resolve")
	flag.Parse()

	addr := fmt.Sprintf("%s:53", dnsAddr)
	fmt.Printf("dailing %s\n\n", addr)
	conn, err := net.Dial("udp", addr)
	handleErr(err)

	defer func() {
		closeErr := conn.Close()
		handleErr(closeErr)
	}()

	req := encodeRequest(resolveAddr)
	m, err := conn.Write(req)
	handleErr(err)
	fmt.Printf("wrote %d bytes\n", m)
	fmt.Printf("%s\n", req)

	res := make([]byte, kbSize)
	n, err := conn.Read(res)
	handleErr(err)
	fmt.Printf("read %d bytes\n", n)
	decodeResponse(res)
}

func handleErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// good blog about bit manipulation in go: https://ops.tips/blog/raw-dns-resolver-in-go/
func encodeRequest(addr string) []byte {
	req := make([]byte, 0, 12)
	req = append(req, []byte{0,1}...) // ID = hardcoded
	// req = append(req, uint(0)) // QR = query
	// req = append(req, []uint{0,0,0,0}...) // Opcode = standard query
	// req = append(req, uint(0)) // AA (server side)
	// req = append(req, uint(0)) // TC = not truncated
	// req = append(req, uint(1)) // RD = recursion is desired
	req = append(req, 1)
	// req = append(req, uint(0)) // RA (server side)
	// req = append(req, []uint{0,0,0}...) // Z (server side)
	// req = append(req, []uint{0,0,0,0}...) // Rcode (server side)
	req = append(req, 0)
	// req = append(req, []uint{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}...) // QDCount = 1 question
	req = append(req, []byte{0,1}...)
	// req = append(req, []uint{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}...) // ANCount (server side)
	// req = append(req, []uint{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}...) // NSCount (server side)
	// req = append(req, []uint{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}...) // ARCount (server side)
	req = append(req, []byte{0,0,0,0,0,0}...)
	return req
}

func decodeResponse(res []byte) {
	fmt.Println("headers")
	fmt.Printf("id: %s\n", res[0:2])
	fmt.Printf("options: %s\n", res[2:3])
	fmt.Printf("server options: %s\n", res[3:4])
	fmt.Printf("question count: %s\n", res[4:6])
	fmt.Printf("answer count: %s\n", res[6:8])
	fmt.Printf("ns count: %s\n", res[8:10])
	fmt.Printf("ar count: %s\n", res[10:12])
	fmt.Println("message")
}
