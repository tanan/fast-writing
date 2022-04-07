package service

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"strings"

	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

type AuthInterceptor struct{}

func NewAuthInterceptor() *AuthInterceptor {
	return &AuthInterceptor{}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)
		userId, err := interceptor.Authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(metadata.AppendToOutgoingContext(ctx, "user_id", userId), req)
	}
}

func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		log.Println("--> stream interceptor: ", info.FullMethod)
		_, err := interceptor.Authorize(ss.Context(), info.FullMethod)
		if err != nil {
			return err
		}
		return handler(srv, ss)
	}
}

func (interceptor *AuthInterceptor) Authorize(ctx context.Context, method string) (string, error) {
	opt := option.WithCredentialsFile("firebase-credentials.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return "", fmt.Errorf("error initializing app: %v", err)
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		return "", fmt.Errorf("error getting Auth client: %v\n", err)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("error getting token from header: %v\n", err)
	}
	fmt.Println(strings.Replace(md.Get("authorization")[0], "Bearer ", "", 1))
	token, err := client.VerifyIDToken(context.Background(), strings.Replace(md.Get("authorization")[0], "Bearer ", "", 1))
	if err != nil {
		return "", fmt.Errorf("error verifying ID token: %v\n", err)
	}
	log.Printf("Verified ID token: %v\n", token)
	return "", nil
}
