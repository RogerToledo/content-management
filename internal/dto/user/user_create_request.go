package user

import "github.com/go/content-management/internal/domain"

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required,min=3,max=30"`
}

func (u *CreateUserRequest) ToDomain() domain.User {
	return domain.User{
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
		Name:     u.Name,
	}
}
