-- name: CreateChirp :one
INSERT INTO chirps (id, created_at, updated_at, body, user_id)
VALUES (
    gen_random_uuid(), now(), now(), $1, $2
)
RETURNING *;

-- name: GetChirps :many
SELECT id, created_at, updated_at, body, user_id
FROM chirps
ORDER BY created_at ASC;

-- name: GetChirp :one
SELECT id, created_at, updated_at, body, user_id
FROM chirps
WHERE id = $1;


