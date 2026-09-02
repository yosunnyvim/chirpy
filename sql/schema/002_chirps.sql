-- +goose Up
create table chirps (
    id UUID primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    body text not null,
    user_id uuid not null references users(id) ON DELETE CASCADE
);

-- +goose down
drop table chirps;
