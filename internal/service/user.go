package service

import (
	"context"

	"github.com/lang/problem/internal/repo/gen"
)

func (s *Service) CreateUser(ctx context.Context, arg gen.CreateUserParams) error {
	return s.repo.CreateUser(ctx, arg)
}

func (s *Service) CheckEmail(ctx context.Context, email string) (gen.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}
