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
