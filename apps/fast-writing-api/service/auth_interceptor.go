package service

import (
	"context"
	"fast-writing-api/config"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"strings"

	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

type AuthInterceptor struct {
	cfg *config.Config
}

func NewAuthInterceptor(cfg *config.Config) *AuthInterceptor {
	return &AuthInterceptor{
		cfg: cfg,
	}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)
		userId, err := interceptor.Authorize(ctx, info.FullMethod)
		if err != nil {
			if err.Error() != "error getting token from header\n" {
				log.Println(err.Error())
			}
			return handler(ctx, req)
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
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("error getting token from header\n")
	}
	var opt option.ClientOption
	if !interceptor.cfg.Application.OnCloud {
		opt = option.WithCredentialsFile("firebase-credentials.json")
	}
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return "", fmt.Errorf("error initializing app: %v", err)
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		return "", fmt.Errorf("error getting Auth client: %v\n", err)
	}
	token, err := client.VerifyIDToken(context.Background(), strings.Replace(md.Get("authorization")[0], "Bearer ", "", 1))
	if err != nil {
		return "", fmt.Errorf("error verifying ID token: %v\n", err)
	}
	return token.UID, nil
}
