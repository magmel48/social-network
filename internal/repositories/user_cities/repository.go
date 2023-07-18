package user_cities

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

func (r *Repository) Create(ctx context.Context, userID, cityID int32) error {
	return r.queries.InsertUserCity(ctx, db.InsertUserCityParams{UserID: userID, CityID: cityID})
}

func (r *Repository) FindByUserID(ctx context.Context, userID int32) (db.UsersCity, error) {
	row, err := r.queries.FindUserCityByUserID(ctx, userID)
	if err != nil {
		return db.UsersCity{}, nil
	}

	return db.UsersCity{UserID: row.UserID, CityID: row.CityID}, nil
}
