package auth

import (
	"context"

	"github.com/lgu-elo/user/pkg/pb"
	"github.com/pkg/errors"
)

type (
	Service struct {
		client Client
		ctx    context.Context
	}
	IService interface {
		GetAllUsers() ([]*User, error)
		GetUserById(id int) (*User, error)
		UpdateUser(user *pb.User) (*User, error)
		DeleteUser(id int64) error
		RegisterUser(user *pb.User) error
	}
)

func NewService(client Client) IService {
	return &Service{client, context.Background()}
}

func (s *Service) GetAllUsers() ([]*User, error) {
	var users []*User

	pbUsers, err := s.client.GetAllUsers(s.ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}

	for _, user := range pbUsers.Users {
		users = append(users, &User{
			ID:    int(user.Id),
			Login: user.Login,
			Role:  user.Role,
			Name:  user.Name,
		})
	}

	return users, nil
}

func (s *Service) GetUserById(id int) (*User, error) {
	pbUser, err := s.client.GetUserById(s.ctx, &pb.UserWithID{
		Id: int64(id),
	})
	if err != nil {
		return nil, err
	}

	return &User{
		ID:    int(pbUser.Id),
		Login: pbUser.Login,
		Name:  pbUser.Name,
		Role:  pbUser.Role,
	}, nil
}

func (s *Service) UpdateUser(user *pb.User) (*User, error) {
	user, err := s.client.UpdateUser(s.ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "cannot update user")
	}

	return &User{
		ID:    int(user.Id),
		Login: user.Login,
		Name:  user.Name,
		Role:  user.Role,
	}, nil
}

func (s *Service) DeleteUser(id int64) error {
	_, err := s.client.DeleteUser(s.ctx, &pb.UserWithID{Id: id})
	if err != nil {
		return errors.Wrap(err, "cannot delete user")
	}
	return nil
}

func (s *Service) RegisterUser(user *pb.User) error {
	_, err := s.client.CreateUser(s.ctx, user)
	if err != nil {
		return errors.Wrap(err, "cannot register user")
	}
	return nil
}
