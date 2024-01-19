-- name: GetOneByID :one
select
    *
from
    "user"
where
    id = $1;

-- name: GetOneByEmail :one
select
    *
from
    "user"
where
    email = $1;

-- name: CreateUser :one
insert into
    "user" (username, email, bio, password_hash, status)
values
    ($1, $2, $3, $4, $5)
returning
    *;

-- name: UpdateUser :one
update "user"
set
    email = coalesce($1, "user".email),
    username = coalesce($2, "user".username),
    password_hash = coalesce($3, "user".password_hash),
    bio = coalesce($4, "user".bio),
    status = coalesce($5, "user".status)
where
    id = $6
returning
    *;