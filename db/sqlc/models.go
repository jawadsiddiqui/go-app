// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"database/sql"
	"time"
)

type Account struct {
	ID       int64     `json:"id"`
	Owner    string    `json:"owner"`
	Balance  int64     `json:"balance"`
	Currency string    `json:"currency"`
	CreateAt time.Time `json:"create_at"`
}

type Author struct {
	ID   int64          `json:"id"`
	Name string         `json:"name"`
	Bio  sql.NullString `json:"bio"`
}
