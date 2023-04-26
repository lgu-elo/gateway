package project

import (
	"context"

	"github.com/lgu-elo/project/pkg/pb"
	"github.com/pkg/errors"
)

type (
	Service struct {
		client Client
		ctx    context.Context
	}
	IService interface {
		GetAllProjects() ([]*Project, error)
		GetProjectById(id int) (*Project, error)
		UpdateProject(project *pb.Project) (*Project, error)
		DeleteProject(id int64) error
		CreateProject(project *pb.Project) error
	}
)

func NewService(client Client) IService {
	return &Service{client, context.Background()}
}

func (s *Service) GetAllProjects() ([]*Project, error) {
	var projects []*Project

	pbProjects, err := s.client.GetAllProjects(s.ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}

	for _, project := range pbProjects.Projects {
		projects = append(projects, &Project{
			ID:          int(project.Id),
			Name:        project.Name,
			Description: project.Description,
		})
	}

	return projects, nil
}

func (s *Service) GetProjectById(id int) (*Project, error) {
	pbProject, err := s.client.GetProjectById(s.ctx, &pb.ProjectWithID{
		Id: int64(id),
	})
	if err != nil {
		return nil, err
	}

	return &Project{
		ID:          int(pbProject.Id),
		Name:        pbProject.Name,
		Description: pbProject.Description,
	}, nil
}

func (s *Service) UpdateProject(project *pb.Project) (*Project, error) {
	pbProject, err := s.client.UpdateProject(s.ctx, project)
	if err != nil {
		return nil, errors.Wrap(err, "cannot update project")
	}

	return &Project{
		ID:          int(pbProject.Id),
		Name:        pbProject.Name,
		Description: pbProject.Description,
	}, nil
}

func (s *Service) DeleteProject(id int64) error {
	_, err := s.client.DeleteProject(s.ctx, &pb.ProjectWithID{Id: id})
	if err != nil {
		return errors.Wrap(err, "cannot delete project")
	}
	return nil
}

func (s *Service) CreateProject(project *pb.Project) error {
	_, err := s.client.CreateProject(s.ctx, project)
	if err != nil {
		return errors.Wrap(err, "cannot delete project")
	}
	return nil
}
