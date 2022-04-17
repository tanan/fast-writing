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
			if err.Error() == "error getting token from header\n" {
				return handler(ctx, req)
			}
			log.Println("auth error: ", err.Error())
			return nil, err

		}
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return handler(ctx, req)
		}
		md.Set("user-id", userId)
		return handler(metadata.NewIncomingContext(context.Background(), md), req)
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
	var app *firebase.App
	var err error
	var idToken string
	if interceptor.cfg.Application.OnCloud {
		app, err = firebase.NewApp(context.Background(), &firebase.Config{
			ProjectID: "anan-project",
		})
		idToken = strings.Replace(md.Get("x-forwarded-authorization")[0], "Bearer ", "", 1)
		log.Println("x-forwarded-authorization:", idToken)
	} else {
		opt := option.WithCredentialsFile("firebase-credentials.json")
		app, err = firebase.NewApp(context.Background(), nil, opt)
		idToken = strings.Replace(md.Get("authorization")[0], "Bearer ", "", 1)
		log.Println("authorization:", idToken)
	}
	if err != nil {
		return "", fmt.Errorf("error initializing app: %v", err)
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		return "", fmt.Errorf("error getting Auth client: %v\n", err)
	}
	token, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return "", err
	}
	return token.UID, nil
}
