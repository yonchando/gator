-- name: CreateFeedFollow :many
INSERT INTO feed_follows (
    id, user_id, feed_id, created_at, updated_at
) VALUES ( $1, $2, $3, $4, $5 ) RETURNING *;

-- name: DeleteFeedFollowByUrlAndUserId :exec
DELETE FROM feed_follows
where user_id = $1 and feed_id = $2;
