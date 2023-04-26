package auth

import (
	"github.com/lgu-elo/gateway/internal/config"
	"github.com/lgu-elo/user/pkg/pb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client pb.UserServiceClient

func NewClient(cfg *config.Cfg) (Client, error) {
	conn, err := grpc.Dial(cfg.Services.User.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrap(err, "canot")
	}

	client := pb.NewUserServiceClient(conn)

	return client, nil
}
