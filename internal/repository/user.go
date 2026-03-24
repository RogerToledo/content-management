package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/go/content-management/apperr"
	"github.com/go/content-management/internal/domain"
	"github.com/go/content-management/internal/infra/db"
	"github.com/go/content-management/internal/models"
	"github.com/go/content-management/internal/pkg/identity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	CreateUser(ctx context.Context, u domain.User) error
	UpdateUser(ctx context.Context, user domain.User) error
	DeleteUser(ctx context.Context, idString string) error
	FindUserByID(ctx context.Context, idString string) (domain.User, error)
	FindAll(ctx context.Context) ([]domain.User, error)
}

type userRepository struct {
	database *pgxpool.Pool
	q        *db.Queries
}

func NewUserRepository(database *pgxpool.Pool) *userRepository {
	return &userRepository{
		database: database,
		q:        db.New(database),}
}

func (r *userRepository) CreateUser(ctx context.Context, u domain.User) error {
	m, err := models.ToUserModel(u)
	if err != nil {
		return err
	}

	m.Active = true

	err = r.q.CreateUser(ctx, db.CreateUserParams{
		ID:       m.Id,
		UserName: m.Username,
		Email:    m.Email,
		Password: m.Password,
		Name:     m.Name,
		Active:   m.Active,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user domain.User) error {
	m, err := models.ToUserModel(user)
	if err != nil {
		return err
	}

	err = r.q.UpdateUser(ctx, db.UpdateUserParams{
		ID:       m.Id,
		Email:    m.Email,
		Name:     m.Name,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, idString string) error {
	id, err := identity.ParseID(idString)
	if err != nil {
		return err
	}

	_, err = r.FindUserByID(ctx, id.String())
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return apperr.MessageError(fmt.Sprintf(apperr.ErrIsNotFound, apperr.UserPT), err)
	}

	err = r.q.DeleteUser(ctx, id)
	if err != nil {
		return apperr.MessageError(apperr.ErrInternalServerError, err)
	}

	return nil
}

func (r *userRepository) FindUserByID(ctx context.Context, idString string) (domain.User, error) {
	id, err := identity.ParseID(idString)
	if err != nil {
		return domain.User{}, err
	}

	row, err := r.q.FindUser(ctx, id)
	if err != nil {
		return domain.User{}, apperr.MessageError(apperr.ErrInternalServerError, err)
	}

	m := models.UserModel{
		Id:       row.ID,
		Username: row.UserName,
		Email:    row.Email,
		Name:     row.Name,
		Active:   row.Active,
	}

	return m.ToDomain(), nil
}

func (r *userRepository) FindAll(ctx context.Context) ([]domain.User, error) {
	rows, err := r.q.FindUsers(ctx)
	if err != nil {
		return []domain.User{}, apperr.MessageError(apperr.ErrInternalServerError, err)
	}
	

	var users []domain.User

	for _, row := range rows {
		users = append(users, domain.User{
			Id:       row.ID.String(),
			Username: row.UserName,
			Email:    row.Email,
			Name:     row.Name,
			Active:   row.Active,
		})
	}

	return users, nil
}
