package routes

import (
	"net/http"

	"github.com/go/content-management/internal/handler"
)

func Setup(
	mux *http.ServeMux,
	userHandler handler.UserHandler,
) {
	//----------- Health -----------
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	//----------- User -----------
	mux.HandleFunc("POST /v1/users", userHandler.CreateUser)
	mux.HandleFunc("PUT /v1/users", userHandler.UpdateUser)
	mux.HandleFunc("DELETE /v1/users/{id}", userHandler.DeleteUser)
	mux.HandleFunc("GET /v1/users/{id}", userHandler.FindUser)
	mux.HandleFunc("GET /v1/users", userHandler.FindUsers)

}
