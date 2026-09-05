-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES (
    gen_random_uuid(), 
    now(),
    now(),
    $1,
    $2
)
RETURNING *;

-- name: GetUserByEmail :one
SELECT id, created_at, updated_at, email, hashed_password
FROM users
where email =$1;

-- name: UpdateUser :one
UPDATE users
SET updated_at = Now(), email = $2, hashed_password = $3
where id = $1
RETURNING *;
