-- +goose Up
CREATE TABLE posts (
    id uuid PRIMARY KEY,
    title varchar(255) not null,
    url varchar(255) unique not null,
    description text,
    published_at timestamp,
    feed_id uuid not null,
    created_at timestamp not null,
    updated_at timestamp not null,

    CONSTRAINT post_feed_id_fk FOREIGN KEY (feed_id)
        REFERENCES feeds(id) on delete cascade
);

-- +goose Down
DROP TABLE posts;
