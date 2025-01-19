-- name: GetUsers :many
SELECT * FROM users;


-- name: CreateUser :one
insert into users (id, name, created_at, updated_at)
values ($1,$2,$3,$4) returning *;

-- name: GetUser :one
select * from users where id = $1 limit 1;

-- name: GetUserByName :one
SELECT * FROM users where name = $1;

-- name: DeleteAllUser :exec
DELETE FROM users;
