-- name: GetUser :one
SELECT user_id, username, created_at 
FROM users
WHERE user_id = $1 LIMIT 1;

-- name: LoginUser :one
SELECT user_id, username, created_at, password 
FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUsers :many
SELECT user_id, username, created_at FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CreateUser :one
INSERT INTO users   ( email, username, password ) 
VALUES              ( $1, $2, $3 )
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1;