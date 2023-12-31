// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: queries.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createUser = `-- name: CreateUser :execresult
INSERT INTO ` + "`" + `users` + "`" + ` (` + "`" + `first_name` + "`" + `, ` + "`" + `last_name` + "`" + `, ` + "`" + `password` + "`" + `, ` + "`" + `gender` + "`" + `, ` + "`" + `birthday` + "`" + `, ` + "`" + `biography` + "`" + `)
VALUES (?, ?, ?, ?, ?, ?)
`

type CreateUserParams struct {
	FirstName string          `db:"first_name" json:"first_name"`
	LastName  string          `db:"last_name" json:"last_name"`
	Password  string          `db:"password" json:"password"`
	Gender    NullUsersGender `db:"gender" json:"gender"`
	Birthday  time.Time       `db:"birthday" json:"birthday"`
	Biography sql.NullString  `db:"biography" json:"biography"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.Password,
		arg.Gender,
		arg.Birthday,
		arg.Biography,
	)
}

const findCityByID = `-- name: FindCityByID :one
SELECT ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + ` FROM ` + "`" + `cities` + "`" + ` WHERE ` + "`" + `id` + "`" + ` = ?
`

type FindCityByIDRow struct {
	ID   int32  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (q *Queries) FindCityByID(ctx context.Context, id int32) (*FindCityByIDRow, error) {
	row := q.db.QueryRowContext(ctx, findCityByID, id)
	var i FindCityByIDRow
	err := row.Scan(&i.ID, &i.Name)
	return &i, err
}

const findCityByName = `-- name: FindCityByName :one
SELECT ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + ` FROM ` + "`" + `cities` + "`" + ` WHERE ` + "`" + `name` + "`" + ` = ?
`

type FindCityByNameRow struct {
	ID   int32  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (q *Queries) FindCityByName(ctx context.Context, name string) (*FindCityByNameRow, error) {
	row := q.db.QueryRowContext(ctx, findCityByName, name)
	var i FindCityByNameRow
	err := row.Scan(&i.ID, &i.Name)
	return &i, err
}

const findUserByID = `-- name: FindUserByID :one
SELECT ` + "`" + `id` + "`" + `, ` + "`" + `password` + "`" + `, ` + "`" + `first_name` + "`" + `, ` + "`" + `last_name` + "`" + `, ` + "`" + `gender` + "`" + `, ` + "`" + `birthday` + "`" + `, ` + "`" + `biography` + "`" + `, ` + "`" + `created_at` + "`" + `
FROM ` + "`" + `users` + "`" + ` WHERE ` + "`" + `id` + "`" + ` = ? LIMIT 1
`

type FindUserByIDRow struct {
	ID        int32           `db:"id" json:"id"`
	Password  string          `db:"password" json:"password"`
	FirstName string          `db:"first_name" json:"first_name"`
	LastName  string          `db:"last_name" json:"last_name"`
	Gender    NullUsersGender `db:"gender" json:"gender"`
	Birthday  time.Time       `db:"birthday" json:"birthday"`
	Biography sql.NullString  `db:"biography" json:"biography"`
	CreatedAt time.Time       `db:"created_at" json:"created_at"`
}

func (q *Queries) FindUserByID(ctx context.Context, id int32) (*FindUserByIDRow, error) {
	row := q.db.QueryRowContext(ctx, findUserByID, id)
	var i FindUserByIDRow
	err := row.Scan(
		&i.ID,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.Gender,
		&i.Birthday,
		&i.Biography,
		&i.CreatedAt,
	)
	return &i, err
}

const findUserCityByUserID = `-- name: FindUserCityByUserID :one
SELECT ` + "`" + `id` + "`" + `, ` + "`" + `city_id` + "`" + `, ` + "`" + `user_id` + "`" + ` FROM ` + "`" + `users_cities` + "`" + ` WHERE ` + "`" + `user_id` + "`" + ` = ?
`

type FindUserCityByUserIDRow struct {
	ID     int32 `db:"id" json:"id"`
	CityID int32 `db:"city_id" json:"city_id"`
	UserID int32 `db:"user_id" json:"user_id"`
}

func (q *Queries) FindUserCityByUserID(ctx context.Context, userID int32) (*FindUserCityByUserIDRow, error) {
	row := q.db.QueryRowContext(ctx, findUserCityByUserID, userID)
	var i FindUserCityByUserIDRow
	err := row.Scan(&i.ID, &i.CityID, &i.UserID)
	return &i, err
}

const insertUserCity = `-- name: InsertUserCity :exec
INSERT INTO ` + "`" + `users_cities` + "`" + ` (` + "`" + `user_id` + "`" + `, ` + "`" + `city_id` + "`" + `) VALUES (?, ?)
`

type InsertUserCityParams struct {
	UserID int32 `db:"user_id" json:"user_id"`
	CityID int32 `db:"city_id" json:"city_id"`
}

func (q *Queries) InsertUserCity(ctx context.Context, arg InsertUserCityParams) error {
	_, err := q.db.ExecContext(ctx, insertUserCity, arg.UserID, arg.CityID)
	return err
}

const upsertCity = `-- name: UpsertCity :execresult
INSERT INTO ` + "`" + `cities` + "`" + ` (` + "`" + `name` + "`" + `) VALUES (?) ON DUPLICATE KEY UPDATE ` + "`" + `name` + "`" + ` = ` + "`" + `name` + "`" + `
`

func (q *Queries) UpsertCity(ctx context.Context, name string) (sql.Result, error) {
	return q.db.ExecContext(ctx, upsertCity, name)
}
