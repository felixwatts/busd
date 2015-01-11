package main

import (
	"flag"
	"fmt"
	"github.com/felixwatts/bus"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", ":8888", "Address to listen on e.g. ':8888'")
}

func main() {
	_, err := bus.Serve(addr)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Scanln()
}
