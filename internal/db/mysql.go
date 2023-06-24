package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func Open(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute)
	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(30)

	return db, err
}
