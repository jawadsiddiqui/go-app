// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: account.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
  owner,
  balance,
  currency
) VALUES (
  $1, $2, $3
)
RETURNING id, owner, balance, currency, create_at
`

type CreateAccountParams struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.queryRow(ctx, q.createAccountStmt, createAccount, arg.Owner, arg.Balance, arg.Currency)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreateAt,
	)
	return i, err
}
