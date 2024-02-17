// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package gen

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createProblemStmt, err = db.PrepareContext(ctx, createProblem); err != nil {
		return nil, fmt.Errorf("error preparing query CreateProblem: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteProblemStmt, err = db.PrepareContext(ctx, deleteProblem); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteProblem: %w", err)
	}
	if q.getAllProblemsStmt, err = db.PrepareContext(ctx, getAllProblems); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllProblems: %w", err)
	}
	if q.getProblemsByCreated_byStmt, err = db.PrepareContext(ctx, getProblemsByCreated_by); err != nil {
		return nil, fmt.Errorf("error preparing query GetProblemsByCreated_by: %w", err)
	}
	if q.getProblemsByIdStmt, err = db.PrepareContext(ctx, getProblemsById); err != nil {
		return nil, fmt.Errorf("error preparing query GetProblemsById: %w", err)
	}
	if q.getUserByEmailStmt, err = db.PrepareContext(ctx, getUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByEmail: %w", err)
	}
	if q.updateProblemStmt, err = db.PrepareContext(ctx, updateProblem); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateProblem: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createProblemStmt != nil {
		if cerr := q.createProblemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createProblemStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteProblemStmt != nil {
		if cerr := q.deleteProblemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteProblemStmt: %w", cerr)
		}
	}
	if q.getAllProblemsStmt != nil {
		if cerr := q.getAllProblemsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllProblemsStmt: %w", cerr)
		}
	}
	if q.getProblemsByCreated_byStmt != nil {
		if cerr := q.getProblemsByCreated_byStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProblemsByCreated_byStmt: %w", cerr)
		}
	}
	if q.getProblemsByIdStmt != nil {
		if cerr := q.getProblemsByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProblemsByIdStmt: %w", cerr)
		}
	}
	if q.getUserByEmailStmt != nil {
		if cerr := q.getUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByEmailStmt: %w", cerr)
		}
	}
	if q.updateProblemStmt != nil {
		if cerr := q.updateProblemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateProblemStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                          DBTX
	tx                          *sql.Tx
	createProblemStmt           *sql.Stmt
	createUserStmt              *sql.Stmt
	deleteProblemStmt           *sql.Stmt
	getAllProblemsStmt          *sql.Stmt
	getProblemsByCreated_byStmt *sql.Stmt
	getProblemsByIdStmt         *sql.Stmt
	getUserByEmailStmt          *sql.Stmt
	updateProblemStmt           *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                          tx,
		tx:                          tx,
		createProblemStmt:           q.createProblemStmt,
		createUserStmt:              q.createUserStmt,
		deleteProblemStmt:           q.deleteProblemStmt,
		getAllProblemsStmt:          q.getAllProblemsStmt,
		getProblemsByCreated_byStmt: q.getProblemsByCreated_byStmt,
		getProblemsByIdStmt:         q.getProblemsByIdStmt,
		getUserByEmailStmt:          q.getUserByEmailStmt,
		updateProblemStmt:           q.updateProblemStmt,
	}
}
