package users

import (
	"context"
	"github.com/magmel48/social-network/internal/db"
)

type Repository struct {
	queries *db.Queries
}

func New(queries *db.Queries) *Repository {
	return &Repository{queries: queries}
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
		Biography: row.Biography,
		CreatedAt: row.CreatedAt,
	}

	return result, err
}
