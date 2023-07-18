package cities

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

func (r *Repository) UpsertCity(ctx context.Context, name string) (db.City, error) {
	sqlResult, err := r.queries.UpsertCity(ctx, name)
	if err != nil {
		return db.City{}, err
	}

	id, err := sqlResult.LastInsertId()
	if err != nil {
		return db.City{}, err
	}

	// update received instead of insert
	if id == 0 {
		row, err := r.queries.FindCityByName(ctx, name)
		if err != nil {
			return db.City{}, err
		}

		id = int64(row.ID)
	}

	return db.City{ID: int32(id), Name: name}, nil
}

func (r *Repository) FindByID(ctx context.Context, id int32) (db.City, error) {
	row, err := r.queries.FindCityByID(ctx, id)
	if err != nil {
		return db.City{}, err
	}

	return db.City{ID: row.ID, Name: row.Name}, nil
}
