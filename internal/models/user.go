package models

import (
	"github.com/go/content-management/internal/domain"
	"github.com/go/content-management/internal/pkg/identity"
	"github.com/google/uuid"
)

type UserModel struct {
	Id       uuid.UUID `db:"id"`
	Username string      `db:"user_name"`
	Password string      `db:"password"`
	Email    string      `db:"email"`
	Name     string      `db:"name"`
	Active   bool        `db:"active"`
}

func (u *UserModel) ToDomain() domain.User {
	return domain.User{
		Id:       u.Id.String(),
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
		Name:     u.Name,
		Active:   u.Active,
	}
}

func ToUserModel(d domain.User) (UserModel, error) {
	id, err := identity.ParseID(d.Id)
	if err != nil {
		return UserModel{}, err
	}

	return UserModel{
		Id:       id,
		Username: d.Username,
		Password: d.Password,
		Email:    d.Email,
		Name:     d.Name,
		Active:   d.Active,
	}, nil
}
