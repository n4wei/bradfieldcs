package main

import (
"fmt"
"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./net.cap")
	if err != nil {
		panic(err)
	}

	pcapHeader := data[:24]
	fmt.Println(pcapHeader)

	// frames := data[24:]
	// fmt.Println(frames)
}
