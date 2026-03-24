package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/go/content-management/apperr"
	dto "github.com/go/content-management/internal/dto/user"
	"github.com/go/content-management/internal/rest"
	"github.com/go/content-management/internal/service"
)

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	FindUser(w http.ResponseWriter, r *http.Request)
	FindUsers(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	service service.UserService
}

func NewUserHandler(svc service.UserService) *userHandler {
	return &userHandler{service: svc}
}

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := rest.Decode[dto.CreateUserRequest](r)
	if err != nil {
		rest.Error(w, http.StatusUnprocessableEntity, apperr.ErrDecoding, err)
		return
	}

	validate := validator.New()

	if err := validate.Struct(user); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			report := make(map[string]string)

			for _, e := range validationErrors {
				report[e.Field()] = apperr.ErrValidation + e.Tag()
			}
			rest.Error(w, http.StatusBadRequest, apperr.ErrBadRequest, report)
			return
		}
		rest.Error(w, http.StatusInternalServerError, apperr.ErrInternalServerError, err.Error())
		return
	}

	u := user.ToDomain()

	err = h.service.CreateUser(ctx, u)
	if err != nil {
		rest.Error(w, http.StatusInternalServerError, fmt.Sprintf(apperr.ErrCreate, apperr.UserPT), err.Error())
		return
	}

	rest.Send(w, fmt.Sprintf(apperr.OkCreate, apperr.UserPT), http.StatusCreated)
}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := rest.Decode[dto.UpdateUserRequest](r)
	if err != nil {
		rest.Error(w, http.StatusUnprocessableEntity, apperr.ErrDecoding, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			report := make(map[string]string)

			for _, e := range validationErrors {
				report[e.Field()] = apperr.ErrValidation + e.Tag()
			}
			rest.Error(w, http.StatusBadRequest, apperr.ErrBadRequest, report)
			return
		}
		rest.Error(w, http.StatusInternalServerError, apperr.ErrInternalServerError, err.Error())
		return
	}

	u := user.ToDomain()

	err = h.service.UpdateUser(ctx, u)
	if err != nil {
		rest.Error(w, http.StatusInternalServerError, apperr.ErrInternalServerError, err.Error())
		return
	}

	rest.Send(w, fmt.Sprintf(apperr.OkUpdate, apperr.UserPT), http.StatusOK)
}

func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idRequest := r.PathValue("id")

	validate := validator.New()

	if err := validate.Var(idRequest, "required,uuid"); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			report := make(map[string]string)

			for _, e := range validationErrors {
				report["id"] = apperr.ErrValidation + e.Tag()
			}
			rest.Error(w, http.StatusBadRequest, apperr.ErrBadRequest, report)
			return
		}
		rest.Error(w, http.StatusInternalServerError, apperr.ErrInternalServerError, err.Error())
		return
	}

	err := h.service.DeleteUser(ctx, idRequest)
	if err != nil {
		rest.Error(w, http.StatusInternalServerError, apperr.ErrInternalServerError, err.Error())
		return
	}

	rest.Send(w, fmt.Sprintf(apperr.OkDelete, apperr.UserPT), http.StatusOK)
}

func (h *userHandler) FindUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idRequest := r.PathValue("id")

	validate := validator.New()

	if err := validate.Var(idRequest, "required,uuid"); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			report := make(map[string]string)

			for _, e := range validationErrors {
				report["id"] = apperr.ErrValidation + e.Tag()
			}
			rest.Error(w, http.StatusBadRequest, apperr.ErrBadRequest, report)
			return
		}
		rest.Error(w, http.StatusInternalServerError, apperr.ErrInternalServerError, err.Error())
		return
	}

	u, err := h.service.FindUserByID(ctx, idRequest)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		rest.Send(w, apperr.ErrDoNotexist, http.StatusOK)
		return
	}

	if err != nil {
		rest.Error(w, http.StatusInternalServerError, apperr.ErrInternalServerError, err.Error())
		return
	}

	rest.Send(w, u, http.StatusOK)
}

func (h *userHandler) FindUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	u, err := h.service.FindAllUsers(ctx)
	if err != nil {
		rest.Error(w, http.StatusInternalServerError, apperr.ErrInternalServerError, err.Error())
		return
	}

	rest.Send(w, u, http.StatusOK)
}
