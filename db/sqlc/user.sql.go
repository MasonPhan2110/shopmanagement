// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "user" (
  username,
  hashed_password,
  full_name,
  email,
  role
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING id, username, hashed_password, full_name, email, role, update_at, created_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	Role           string `json:"role"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
		arg.Role,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Role,
		&i.UpdateAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, hashed_password, full_name, email, role, update_at, created_at FROM "user"
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Role,
		&i.UpdateAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUserForUpdate = `-- name: GetUserForUpdate :one
SELECT id, username, hashed_password, full_name, email, role, update_at, created_at FROM "user"
WHERE username = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetUserForUpdate(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserForUpdate, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Role,
		&i.UpdateAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUserHashedPassword = `-- name: UpdateUserHashedPassword :one
UPDATE "user"
SET hashed_password = $2, update_at = (now())
WHERE id = $1
RETURNING id, username, hashed_password, full_name, email, role, update_at, created_at
`

type UpdateUserHashedPasswordParams struct {
	ID             int64  `json:"id"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) UpdateUserHashedPassword(ctx context.Context, arg UpdateUserHashedPasswordParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserHashedPassword, arg.ID, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Role,
		&i.UpdateAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUserRole = `-- name: UpdateUserRole :one
UPDATE "user"
SET role = $2, update_at = (now())
WHERE id = $1
RETURNING id, username, hashed_password, full_name, email, role, update_at, created_at
`

type UpdateUserRoleParams struct {
	ID   int64  `json:"id"`
	Role string `json:"role"`
}

func (q *Queries) UpdateUserRole(ctx context.Context, arg UpdateUserRoleParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserRole, arg.ID, arg.Role)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Role,
		&i.UpdateAt,
		&i.CreatedAt,
	)
	return i, err
}