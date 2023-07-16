package users

import (
	"context"
	"database/sql"
	"github.com/magmel48/social-network/internal/db"
)

type Repository struct {
	queries  *db.Queries
	database *sql.DB
}

func New(queries *db.Queries, database *sql.DB) *Repository {
	return &Repository{queries: queries, database: database}
}

func (r *Repository) FindByID(ctx context.Context, id int32) (db.User, error) {
	row, err := r.queries.FindUserByID(ctx, id)
	if err != nil {
		return db.User{}, err
	}

	result := db.User{
		ID:        id,
		FirstName: row.FirstName,
		LastName:  row.LastName,
		Gender:    row.Gender,
		Birthday:  row.Birthday,
		CreatedAt: row.CreatedAt,
	}

	return result, err
}
