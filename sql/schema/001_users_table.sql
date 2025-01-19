-- +goose Up
CREATE TABLE users (
    id uuid PRIMARY KEY,
    name varchar(255) unique not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

-- +goose Down
DROP TABLE users;
