package auth

import (
	"context"

	"github.com/lgu-elo/auth/pkg/pb"
	"github.com/pkg/errors"
)

type (
	Service struct {
		client Client
		ctx    context.Context
	}
	IService interface {
		CreateNewUser(credentials FullCredentials) (string, error)
		Login(credentials Credentials) (string, error)
	}
)

func NewService(client Client) IService {
	return &Service{client, context.Background()}
}

func (s *Service) CreateNewUser(creds FullCredentials) (string, error) {
	token, err := s.client.Register(s.ctx, &pb.BasicCredentials{
		Username: creds.Login,
		Password: creds.Password,
		Name:     creds.Name,
	})
	if err != nil {
		return "", errors.Wrap(err, "cannot register user")
	}

	return token.Token, nil
}

func (s *Service) Login(credentials Credentials) (string, error) {
	token, err := s.client.Login(s.ctx, &pb.Credentials{
		Username: credentials.Login,
		Password: credentials.Password,
	})
	if err != nil {
		return "", errors.Wrap(err, "cannot login user")
	}

	return token.Token, nil
}
