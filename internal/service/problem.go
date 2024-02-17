package service

import (
	"context"

	"github.com/lang/problem/internal/repo/gen"
)

type Servicer interface {
	CreateProblem(ctx context.Context, arg gen.CreateProblemParams) error
	CreateUser(ctx context.Context, arg gen.CreateUserParams) error
	DeleteProblem(ctx context.Context, arg gen.DeleteProblemParams) error
	GetAllProblems(ctx context.Context) ([]gen.Problem, error)
	GetProblemsById(ctx context.Context, id int32) (gen.Problem, error)
	CheckEmail(ctx context.Context, email string) (gen.User, error)
	UpdateProblem(ctx context.Context, arg gen.UpdateProblemParams) (gen.UpdateProblemRow, error)
}

var _ Servicer = (*Service)(nil)

type Service struct {
	repo gen.Querier
}

func NewService(repo gen.Querier) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateProblem(ctx context.Context, arg gen.CreateProblemParams) error {
	return s.repo.CreateProblem(ctx, arg)
}

func (s *Service) GetAllProblems(ctx context.Context) ([]gen.Problem, error) {
	return s.repo.GetAllProblems(ctx)
}

func (s *Service) GetProblemsById(ctx context.Context, id int32) (gen.Problem, error) {
	return s.repo.GetProblemsById(ctx, id)
}

func (s *Service) UpdateProblem(ctx context.Context, arg gen.UpdateProblemParams) (gen.UpdateProblemRow, error) {
	return s.repo.UpdateProblem(ctx, arg)
}

func (s *Service) DeleteProblem(ctx context.Context, arg gen.DeleteProblemParams) error {
	return s.repo.DeleteProblem(ctx, arg)
}
