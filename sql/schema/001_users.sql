-- +goose Up
CREATE TABLE users (
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    name varchar(255) unique not null
);

-- +goose Down
DROP TABLE users;