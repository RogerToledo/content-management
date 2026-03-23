-- name: CreateUser :exec
INSERT INTO account (id, user_name, email, password, name, active) VALUES ($1,$2,$3,$4, $5, $6)