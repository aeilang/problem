// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: problem.sql

package gen

import (
	"context"

	"github.com/google/uuid"
)

const createProblem = `-- name: CreateProblem :exec
insert into problems (title, description, is_done, created_by)
values ($1, $2, $3, $4)
`

type CreateProblemParams struct {
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	IsDone      bool      `db:"is_done" json:"is_done"`
	CreatedBy   uuid.UUID `db:"created_by" json:"created_by"`
}

func (q *Queries) CreateProblem(ctx context.Context, arg CreateProblemParams) error {
	_, err := q.exec(ctx, q.createProblemStmt, createProblem,
		arg.Title,
		arg.Description,
		arg.IsDone,
		arg.CreatedBy,
	)
	return err
}

const deleteProblem = `-- name: DeleteProblem :exec
delete from problems where created_by = $1 and id = $2
`

type DeleteProblemParams struct {
	CreatedBy uuid.UUID `db:"created_by" json:"created_by"`
	ID        int32     `db:"id" json:"id"`
}

func (q *Queries) DeleteProblem(ctx context.Context, arg DeleteProblemParams) error {
	_, err := q.exec(ctx, q.deleteProblemStmt, deleteProblem, arg.CreatedBy, arg.ID)
	return err
}

const getAllProblems = `-- name: GetAllProblems :many
select id, title, description, is_done, created_by from problems
`

func (q *Queries) GetAllProblems(ctx context.Context) ([]Problem, error) {
	rows, err := q.query(ctx, q.getAllProblemsStmt, getAllProblems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Problem
	for rows.Next() {
		var i Problem
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.IsDone,
			&i.CreatedBy,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProblemsByCreated_by = `-- name: GetProblemsByCreated_by :many
select id, title, description, is_done, created_by from problems
where created_by = $1
`

func (q *Queries) GetProblemsByCreated_by(ctx context.Context, createdBy uuid.UUID) ([]Problem, error) {
	rows, err := q.query(ctx, q.getProblemsByCreated_byStmt, getProblemsByCreated_by, createdBy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Problem
	for rows.Next() {
		var i Problem
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.IsDone,
			&i.CreatedBy,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProblemsById = `-- name: GetProblemsById :one
select id, title, description, is_done, created_by from problems
where id = $1
`

func (q *Queries) GetProblemsById(ctx context.Context, id int32) (Problem, error) {
	row := q.queryRow(ctx, q.getProblemsByIdStmt, getProblemsById, id)
	var i Problem
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.IsDone,
		&i.CreatedBy,
	)
	return i, err
}

const updateProblem = `-- name: UpdateProblem :one
update problems set title = $1, description = $2
where created_by = $3 and id = $4 returning id, title, description, is_done
`

type UpdateProblemParams struct {
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	CreatedBy   uuid.UUID `db:"created_by" json:"created_by"`
	ID          int32     `db:"id" json:"id"`
}

type UpdateProblemRow struct {
	ID          int32  `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	IsDone      bool   `db:"is_done" json:"is_done"`
}

func (q *Queries) UpdateProblem(ctx context.Context, arg UpdateProblemParams) (UpdateProblemRow, error) {
	row := q.queryRow(ctx, q.updateProblemStmt, updateProblem,
		arg.Title,
		arg.Description,
		arg.CreatedBy,
		arg.ID,
	)
	var i UpdateProblemRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.IsDone,
	)
	return i, err
}
