-- name: CreateUser :exec
INSERT INTO account (id, user_name, email, password, name, active) VALUES (@id, @user_name, @email, @password, @name, @active);

-- name: UpdateUser :exec
UPDATE account 
SET email = @email,
    "name" = @name
WHERE id = @id;

-- name: DeleteUser :exec
DELETE FROM account WHERE id = @id;

-- name: FindUser :one
SELECT
    id,
    user_name,
    email,
    name,
    active
FROM account 
WHERE id = @id;

-- name: FindUsers :many
SELECT
    id,
    user_name,
    email,
    name,
    active
FROM account;