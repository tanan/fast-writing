package client

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/idtoken"
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
		idTokenSource, err := idtoken.NewTokenSource(context.Background(), "https://search-api-cmfot4feta-an.a.run.app")
		if err != nil {
			return nil, errors.New("cannot get idTokenSource:" + err.Error())
		}
		dialOptions = append(dialOptions, grpc.WithPerRPCCredentials(oauth.TokenSource{
			TokenSource: idTokenSource,
		}))
	}

	addr := fmt.Sprintf("%s:%s", hostname, port)
	return grpc.Dial(addr, dialOptions...)
}
