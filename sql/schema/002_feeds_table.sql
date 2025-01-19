-- +goose Up
create table feeds
(
    id         uuid           not null
        constraint feeds_pk
            primary key,
    name       varchar(255) not null,
    url        varchar(255) unique not null,
    user_id    uuid           not null
        constraint feeds_fk_user_id references users (id) on delete cascade,
    created_at timestamp      not null,
    updated_at timestamp      not null
);

-- +goose Down
DROP TABLE feeds;
