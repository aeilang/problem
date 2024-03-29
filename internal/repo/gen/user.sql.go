// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package gen

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :exec
insert into users(id, username, email, password_hash)
values ($1, $2, $3, $4)
`

type CreateUserParams struct {
	ID           uuid.UUID `db:"id" json:"id"`
	Username     string    `db:"username" json:"username"`
	Email        string    `db:"email" json:"email"`
	PasswordHash string    `db:"password_hash" json:"password_hash"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.exec(ctx, q.createUserStmt, createUser,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.PasswordHash,
	)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
select id, username, email, password_hash from users where email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.queryRow(ctx, q.getUserByEmailStmt, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
	)
	return i, err
}
