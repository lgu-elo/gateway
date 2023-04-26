// package task

// import (
// 	"github.com/lgu-elo/gateway/internal/config"
// 	"github.com/lgu-elo/task/pkg/pb"
// 	"github.com/pkg/errors"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// type Client pb.TaskServiceClient

// func NewClient(cfg *config.Cfg) (Client, error) {
// 	conn, err := grpc.Dial(cfg.Services.Task.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		return nil, errors.Wrap(err, "canot")
// 	}

// 	client := pb.NewTaskServiceClient(conn)

// 	return client, nil
// }
