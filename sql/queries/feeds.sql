-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedWithUser :many
SELECT * FROM feeds
left join users on users.id = feeds.user_id
;

-- name: CreateFeed :one
insert into feeds (id, name, url, user_id, created_at, updated_at)
values ($1,$2,$3,$4,$5,$6) returning *;

-- name: GetFeed :one
select * from feeds where id = $1 limit 1;

-- name: GetFeedByName :one
SELECT * FROM feeds where name = $1;

-- name: GetFeedByUrl :one
SELECT * FROM feeds where url = $1;

-- name: DeleteAllFeed :exec
DELETE FROM feeds;

