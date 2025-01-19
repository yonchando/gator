-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedWithUser :many
SELECT f.id as feed_id, f.name as feed_name, f.url,
    u.id as user_id, u.name as user_name
FROM feeds f
left join users u on u.id = f.user_id;

-- name: GetFeedFollowsForUser :many
SELECT f.id as feed_id, f.name as feed_name, 
    u.id as user_id, u.name as user_name
    FROM feeds f
    left join feed_follows ff on f.id = ff.feed_id
    left join users u on ff.user_id = u.id
    WHERE u.id = $1;

-- name: GetFeed :one
select * from feeds where id = $1 limit 1;

-- name: GetFeedByName :one
SELECT * FROM feeds where name = $1;

-- name: GetFeedByUrl :one
SELECT * FROM feeds where url = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds order by updated_at asc, last_fetched_at desc;

-- name: CreateFeed :one
insert into feeds (id, name, url, user_id, created_at, updated_at)
values ($1,$2,$3,$4,$5,$6) returning *;

-- name: MarkFeedFetched :exec
UPDATE feeds SET last_fetched_at = $1, updated_at = $2 WHERE id = $3;

-- name: DeleteAllFeed :exec
DELETE FROM feeds;
