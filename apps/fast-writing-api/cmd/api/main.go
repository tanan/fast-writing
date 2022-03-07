package main

import (
	"fast-writing-api/config"
	"fast-writing-api/database"
	"fast-writing-api/pb"
	"fast-writing-api/service"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var (
	path = flag.String("c", "./", "config directory path")
	port = flag.Int("p", 10001, "server running port")
)

func main() {
	flag.Parse()
	// 起動するポート番号を指定しています。
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// gRPCサーバーの生成
	s := grpc.NewServer()
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
	pb.RegisterUserServiceServer(s, service.NewUserService(sqlHandler))
	pb.RegisterWritingServiceServer(s, service.NewWritingService(sqlHandler))

	// サーバーリフレクションを有効にしています。
	// 有効にすることでシリアライズせずとも後述する`grpc_cli`で動作確認ができるようになります。
	reflection.Register(s)
	// サーバーを起動
	if err := s.Serve(listenPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
