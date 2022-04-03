package main

import (
	"fast-writing-api/client"
	"fast-writing-api/config"
	"fast-writing-api/database"
	"fast-writing-api/service"
	"fast-writing/pkg/pb"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
)

var (
	path = flag.String("c", "./", "config directory path")
)

func main() {
	flag.Parse()

	interceptor := service.NewAuthInterceptor()
	// gRPCサーバーの生成
	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)
	// 自動生成された関数に、サーバと実際に処理を行うメソッドを実装したハンドラを設定します。
	// protoファイルで定義した`RockPaperScissorsService`に対応しています。
	c, err := config.LoadConfig(*path)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	sqlHandler, err := database.NewSQLHandler(c.GetSQLConnection(), c.Database.Debug)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	defer sqlHandler.Close()

	//// for search
	cc, err := client.NewClientConn(c.Search.Host, strconv.Itoa(c.Search.Port), c.Search.Ssl, c.Search.Token)
	if err != nil {
		log.Fatalf("failed to client connection: %v", err)
	}
	searchClient := pb.NewSearchServiceClient(cc)
	pb.RegisterSearchServiceServer(s, service.NewSearchService(sqlHandler, searchClient))
	pb.RegisterUserServiceServer(s, service.NewUserService(sqlHandler))
	pb.RegisterWritingServiceServer(s, service.NewWritingService(sqlHandler, searchClient))

	// サーバーリフレクションを有効にしています。
	// 有効にすることでシリアライズせずとも後述する`grpc_cli`で動作確認ができるようになります。
	reflection.Register(s)

	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Application.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := s.Serve(listenPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
