// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: counter.sql

package repository

import (
	"context"
)

const decrementCounter = `-- name: DecrementCounter :exec
UPDATE counter SET count = count - 1 WHERE id = 1
`

func (q *Queries) DecrementCounter(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, decrementCounter)
	return err
}

const getCounter = `-- name: GetCounter :one
SELECT count FROM counter WHERE id = 1
`

func (q *Queries) GetCounter(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getCounter)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const incrementCounter = `-- name: IncrementCounter :exec
UPDATE counter SET count = count + 1 WHERE id = 1
`

func (q *Queries) IncrementCounter(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, incrementCounter)
	return err
}

const initializeCounter = `-- name: InitializeCounter :exec
INSERT INTO counter (id, count) VALUES (1, 0) ON CONFLICT (id) DO NOTHING
`

func (q *Queries) InitializeCounter(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, initializeCounter)
	return err
}
