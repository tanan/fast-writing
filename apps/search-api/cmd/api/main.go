package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	pb "fast-writing/pkg/pb"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"search-api/config"
	"search-api/service"
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

	s := grpc.NewServer()
	_, err = config.LoadConfig(*path)
	if err != nil{
		log.Fatalf("failed to serve: %v", err)
	}

	pb.RegisterSearchServiceServer(s, service.NewSearchService())

	reflection.Register(s)

	if err := s.Serve(listenPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}