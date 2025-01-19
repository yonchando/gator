-- name: CreatePost :one
INSERT INTO posts (id, title, url, description, published_at, feed_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;


-- name: GetPosts :many
SELECT *
FROM posts
order by updated_at desc
limit $1;

-- name: DeleteAllPosts :exec
DELETE FROM posts;
