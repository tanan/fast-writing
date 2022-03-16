package client

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc/credentials/oauth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewClientConn(hostname string, port string, useSSL bool, useToken bool) (*grpc.ClientConn, error) {
	var do grpc.DialOption
	if useSSL {
		tCreds := credentials.NewClientTLSFromCert(nil, "")
		do = grpc.WithTransportCredentials(tCreds)
	} else {
		do = grpc.WithInsecure()
	}

	dialOptions := []grpc.DialOption{do}

	if useToken {
		tokenSource, err := google.DefaultTokenSource(context.Background(), "https://www.googleapis.com/auth/cloud-platform")
		if err != nil {
			return nil, errors.New("cannot get tokenSource:" + err.Error())
		}
		dialOptions = append(dialOptions, grpc.WithPerRPCCredentials(oauth.TokenSource{
			TokenSource: tokenSource,
		}))
	}

	addr := fmt.Sprintf("%s:%s", hostname, port)
	return grpc.Dial(addr, dialOptions...)
}
