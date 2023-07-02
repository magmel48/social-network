// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error)
	FindUserByID(ctx context.Context, id int32) (*FindUserByIDRow, error)
	FindUserWithCheckingPassword(ctx context.Context, arg FindUserWithCheckingPasswordParams) (*FindUserWithCheckingPasswordRow, error)
}

var _ Querier = (*Queries)(nil)
