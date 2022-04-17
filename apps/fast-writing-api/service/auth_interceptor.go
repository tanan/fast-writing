package service

import (
	"context"
	"errors"
	"fast-writing-api/config"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"regexp"
	"strings"
)

type AuthInterceptor struct {
	cfg *config.Config
}

func NewAuthInterceptor(cfg *config.Config) *AuthInterceptor {
	return &AuthInterceptor{
		cfg: cfg,
	}
}

const (
	MissingAuthorizationHeaderError = "failed to get (authorization|x-forwarded-authorization) on metadata"
)

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)
		userId, err := interceptor.Authorize(ctx, info.FullMethod)
		if err != nil {
			r := regexp.MustCompile(MissingAuthorizationHeaderError)
			if r.MatchString(err.Error()) {
				return handler(ctx, req)
			}
			log.Println("auth error: ", err.Error())
			return nil, status.Error(codes.Unauthenticated, codes.Unauthenticated.String())
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

func (interceptor *AuthInterceptor) getAuthorizationToken(ctx context.Context, key string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	value := md.Get(key)
	if !ok || len(value) == 0 {
		return "", errors.New(fmt.Sprintf("failed to get %v on metadata", key))
	}
	return value[0], nil
}

func (interceptor *AuthInterceptor) Authorize(ctx context.Context, method string) (string, error) {
	var app *firebase.App
	var err error
	var idToken string
	if interceptor.cfg.Application.OnCloud {
		app, err = firebase.NewApp(context.Background(), &firebase.Config{
			ProjectID: "anan-project",
		})
		if token, err := interceptor.getAuthorizationToken(ctx, "x-forwarded-authorization"); err != nil {
			return "", err
		} else {
			idToken = strings.Replace(token, "Bearer ", "", 1)
		}
	} else {
		opt := option.WithCredentialsFile("firebase-credentials.json")
		app, err = firebase.NewApp(context.Background(), nil, opt)
		if token, err := interceptor.getAuthorizationToken(ctx, "authorization"); err != nil {
			return "", err
		} else {
			idToken = strings.Replace(token, "Bearer ", "", 1)
		}
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
