// package project

// import (
// 	"github.com/lgu-elo/gateway/internal/config"
// 	"github.com/lgu-elo/project/pkg/pb"
// 	"github.com/pkg/errors"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// type Client pb.ProjectServiceClient

// func NewClient(cfg *config.Cfg) (Client, error) {
// 	conn, err := grpc.Dial(cfg.Services.Project.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		return nil, errors.Wrap(err, "canot")
// 	}

// 	client := pb.NewProjectServiceClient(conn)

// 	return client, nil
// }
