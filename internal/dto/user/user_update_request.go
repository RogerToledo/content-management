package user

import (
	"github.com/go/content-management/internal/domain"
)

type UpdateUserRequest struct {
	Id    string `json:"id" validate:"required,uuid"`
	Email string `json:"email" validate:"omitempty,email"`
	Name  string `json:"name" validate:"omitempty,min=3,max=30"`
}

func (u *UpdateUserRequest) ToDomain() domain.User {
	return domain.User{
		Email: u.Email,
		Name:  u.Name,
	}
}
