-- +goose Up
CREATE TABLE feed_follows (
    id uuid primary key not null,
    user_id uuid not null,
    feed_id uuid not null,
    created_at timestamp not null,
    updated_at timestamp not null,

    CONSTRAINT feed_follows_feed_fk FOREIGN KEY (feed_id)
        REFERENCES feeds(id),
    CONSTRAINT feed_follows_user_fk FOREIGN KEY (user_id)
        REFERENCES users(id) on delete cascade,
    CONSTRAINT feed_follows_unique unique (user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;
