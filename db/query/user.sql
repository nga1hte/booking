-- name: CreateUser :one
INSERT INTO users (
    full_name,
    email,
    mobile_number,
    password,
    user_type
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY full_name
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
    set full_name = $2,
    email = $3,
    mobile_number = $4,
    password = $5
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;