package service

import (
	"context"

	"github.com/go/content-management/internal/domain"
	"github.com/go/content-management/internal/pkg/identity"
	"github.com/go/content-management/internal/repository"
	"github.com/go/content-management/internal/secutity"
)

type UserService interface {
	CreateUser(ctx context.Context, u domain.User) error
	UpdateUser(ctx context.Context, u domain.User) error
	DeleteUser(ctx context.Context, id string) error
	FindUserByID(ctx context.Context, id string) (domain.User, error)
	FindAllUsers(ctx context.Context) ([]domain.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *userService {
	return &userService{
		repo: r,
	}
}

func (s *userService) CreateUser(ctx context.Context, u domain.User) error {
	hashedPassword, err := secutity.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	id, err := identity.GenerateID()
	if err != nil {
		return err
	}

	u.Id = id.String()

	return s.repo.CreateUser(ctx, u)
}

func (s *userService) UpdateUser(ctx context.Context, u domain.User) error {
	err := s.repo.UpdateUser(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {
	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) FindUserByID(ctx context.Context, id string) (domain.User, error) {
	user, err := s.repo.FindUserByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *userService) FindAllUsers(ctx context.Context) ([]domain.User, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
