package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	path = flag.String("c", "./", "config directory path")
	port = flag.Int("p", 10001, "server running port")
)

func main() {
	flag.Parse()
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

}