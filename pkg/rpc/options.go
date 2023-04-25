package rpc

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Option = func() *grpc.ServerOption

func WithLoggerLogrus(logger *logrus.Logger) Option {
	return func() *grpc.ServerOption {
		logrusEntry := logrus.NewEntry(logger)
		grpc_logrus.ReplaceGrpcLogger(logrusEntry)
		optsLogrus := []grpc_logrus.Option{
			grpc_logrus.WithLevels(grpc_logrus.DefaultCodeToLevel),
		}

		opt := grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logrusEntry, optsLogrus...),
		)

		return &opt
	}
}
