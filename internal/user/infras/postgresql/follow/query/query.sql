-- name: AllByFollowedUserID :many
select
    id,
    followed_user_id,
    following_user_id,
    status,
    created_date,
    updated_date
from
    follow
where
    followed_user_id = $1
    and status = ANY ($2::text[])
    and id < $3
order by
    id desc
limit
    $4;

-- name: CreateFollow :one
insert into
    follow (followed_user_id, following_user_id)
values
    ($1, $2)
returning
    *;

-- name: UpdateFollow :one
update follow
set
    status = coalesce($1, "follow".status)
where
    followed_user_id = $2
    and following_user_id = $3
returning
    *;

-- name: CountByFollowedUserID :one
select
    count(*)
from
    follow
where
    followed_user_id = $1;

-- name: GetOne :one
select
    *
from
    follow
where
    followed_user_id = $1
    and following_user_id = $2;