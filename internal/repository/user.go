package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/go/content-management/apperr"
	"github.com/go/content-management/internal/domain"
	"github.com/go/content-management/internal/models"
	"github.com/go/content-management/internal/pkg/identity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) error
	UpdateUser(ctx context.Context, user domain.User) error
	DeleteUser(ctx context.Context, id string) error
	FindUserByID(ctx context.Context, id string) (domain.User, error)
	FindAll(ctx context.Context) ([]domain.User, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(ctx context.Context, u domain.User) error {
	m, err := models.ToUserModel(u)
	if err != nil {
		return err
	}

	m.Active = true

	_, err = r.db.Exec(ctx, CreateUserQuery, m.Id, m.Username, m.Email, m.Password, m.Name, m.Active)
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

	_, err = r.db.Exec(ctx, UpdateUserQuery, m.Username, m.Email, m.Password, m.Name, m.Active, m.Id)
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

	_, err = r.db.Exec(ctx, DeleteUserQuery, id)
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

	var m models.UserModel

	err = r.db.QueryRow(ctx, FindUserQuery, id).Scan(
		m.Id, m.Username, m.Email, m.Active, m.Name,
	)
	if err != nil {
		return domain.User{}, apperr.MessageError(apperr.ErrInternalServerError, err)
	}

	return m.ToDomain(), nil
}

func (r *userRepository) FindAll(ctx context.Context) ([]domain.User, error) {
	rows, err := r.db.Query(ctx, FindAllUserQuery)
	if err != nil {
		return []domain.User{}, apperr.MessageError(apperr.ErrInternalServerError, err)
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var m models.UserModel
		err := rows.Scan(
			&m.Id, &m.Username, &m.Email, &m.Name, &m.Active,
		)
		if err != nil {
			return []domain.User{}, apperr.MessageError(apperr.ErrInternalServerError, err)
		}

		users = append(users, m.ToDomain())
	}

	return users, nil
}
