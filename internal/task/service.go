package task

import (
	"context"

	"github.com/lgu-elo/task/pkg/pb"
	"github.com/pkg/errors"
)

type (
	Service struct {
		client Client
		ctx    context.Context
	}
	IService interface {
		GetAllTasks() ([]*Task, error)
		GetTaskById(id int) (*Task, error)
		UpdateTask(task *pb.Task) (*Task, error)
		DeleteTask(id int64) error
		CreateTask(task *pb.Task) error
	}
)

func NewService(client Client) IService {
	return &Service{client, context.Background()}
}

func (s *Service) GetAllTasks() ([]*Task, error) {
	var tasks []*Task

	pbTasks, err := s.client.GetAllTasks(s.ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}

	for _, task := range pbTasks.Tasks {
		tasks = append(tasks, &Task{
			ID:          int(task.Id),
			Name:        task.Name,
			Status:      task.status,
			User_id:     int(task.UserId),
			Project_id:  int(task.ProjectId),
			Description: task.Description,
		})
	}

	return tasks, nil
}

func (s *Service) GetTaskById(id int) (*Task, error) {
	pbTask, err := s.client.GetTaskById(s.ctx, &pb.TaskWithID{
		Id: int64(id),
	})
	if err != nil {
		return nil, err
	}

	return &Task{
		Name:        pbTask.Name,
		Description: pbTask.Description,
		Status:      pbTask.Status,
		User_id:     int(pbTask.UserId),
		Project_id:  int(task.ProjectId),
	}, nil
}

func (s *Service) UpdateTask(task *pb.Task) (*Task, error) {
	pbTask, err := s.client.UpdateTask(s.ctx, task)
	if err != nil {
		return nil, errors.Wrap(err, "cannot update task")
	}

	return &Task{
		ID:          int(pbTask.Id),
		Name:        pbTask.Name,
		Status:      pbTask.Status,
		User_id:     int(pbTask.UserId),
		Project_id:  int(task.ProjectId),
		Description: pbTask.Description,
	}, nil
}

func (s *Service) DeleteTask(id int64) error {
	_, err := s.client.DeleteTask(s.ctx, &pb.TaskWithID{Id: id})
	if err != nil {
		return errors.Wrap(err, "cannot delete task")
	}
	return nil
}

func (s *Service) CreateTask(task *pb.Task) error {
	_, err := s.client.CreateTask(s.ctx, task)
	if err != nil {
		return errors.Wrap(err, "cannot delete task")
	}
	return nil
}
