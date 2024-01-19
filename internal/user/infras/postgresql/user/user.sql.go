// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: user.sql

package postgresql

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
insert into
    "user" (username, email, bio, password_hash, status)
values
    ($1, $2, $3, $4, $5)
returning
    id, username, email, bio, password_hash, status, created_date, updated_date
`

type CreateUserParams struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Bio          string `json:"bio"`
	PasswordHash string `json:"password_hash"`
	Status       string `json:"status"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.Bio,
		arg.PasswordHash,
		arg.Status,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Bio,
		&i.PasswordHash,
		&i.Status,
		&i.CreatedDate,
		&i.UpdatedDate,
	)
	return i, err
}

const getOneByEmail = `-- name: GetOneByEmail :one
select
    id,
    email,
    username,
    password_hash,
    status,
    bio
from
    "user"
where
    email = $1
`

type GetOneByEmailRow struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"password_hash"`
	Status       string    `json:"status"`
	Bio          string    `json:"bio"`
}

func (q *Queries) GetOneByEmail(ctx context.Context, email string) (GetOneByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getOneByEmail, email)
	var i GetOneByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.Status,
		&i.Bio,
	)
	return i, err
}

const getOneByID = `-- name: GetOneByID :one
select
    id,
    username,
    email,
    bio,
    status,
    created_date,
    updated_date
from
    "user"
where
    id = $1
`

type GetOneByIDRow struct {
	ID          uuid.UUID    `json:"id"`
	Username    string       `json:"username"`
	Email       string       `json:"email"`
	Bio         string       `json:"bio"`
	Status      string       `json:"status"`
	CreatedDate time.Time    `json:"created_date"`
	UpdatedDate sql.NullTime `json:"updated_date"`
}

func (q *Queries) GetOneByID(ctx context.Context, id uuid.UUID) (GetOneByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getOneByID, id)
	var i GetOneByIDRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Bio,
		&i.Status,
		&i.CreatedDate,
		&i.UpdatedDate,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
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
    id, username, email, bio, password_hash, status, created_date, updated_date
`

type UpdateUserParams struct {
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"password_hash"`
	Bio          string    `json:"bio"`
	Status       string    `json:"status"`
	ID           uuid.UUID `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.Email,
		arg.Username,
		arg.PasswordHash,
		arg.Bio,
		arg.Status,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Bio,
		&i.PasswordHash,
		&i.Status,
		&i.CreatedDate,
		&i.UpdatedDate,
	)
	return i, err
}
