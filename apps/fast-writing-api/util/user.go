package util

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
)

func GetUserId(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	userId := md.Get("user-id")
	if !ok || len(userId) == 0 {
		return "", errors.New("failed to get user-id on metadata")
	}
	return userId[0], nil
}
