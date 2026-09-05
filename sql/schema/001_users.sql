-- +goose Up
create table users (
    id UUID primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    email text not null unique
);

-- +goose Down
drop table users;

-- +goose Up
ALTER TABLE users
ADD COLUMN is_chirpy_red BOOLEAN NOT NULL DEFAULT FALSE;

-- +goose Down
ALTER TABLE users
DROP COLUMN is_chirpy_red;
