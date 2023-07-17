package users

import (
	"context"
	"database/sql"
	"github.com/magmel48/social-network/internal/db"
	"golang.org/x/crypto/bcrypt"
)

func (r *Repository) Register(ctx context.Context, user db.User, city, biography *string) (db.User, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return db.User{}, err
	}

	tx, err := r.database.BeginTx(ctx, nil)
	if err != nil {
		return db.User{}, err
	}
	defer tx.Rollback()

	q := r.queries.WithTx(tx)

	createUserParams := db.CreateUserParams{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  string(pass),
		Birthday:  user.Birthday,
	}
	if biography != nil {
		createUserParams.Biography = sql.NullString{String: *biography, Valid: true}
	}

	sqlResult, err := q.CreateUser(ctx, createUserParams)
	if err != nil {
		return db.User{}, err
	}

	userID, err := sqlResult.LastInsertId()
	if err != nil {
		return db.User{}, err
	}

	result := db.User{
		ID:        int32(userID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Gender:    user.Gender,
		Birthday:  user.Birthday,
	}

	if city != nil {
		sqlResult, err := q.UpsertCity(ctx, *city)
		if err != nil {
			return db.User{}, err
		}

		cityID, err := sqlResult.LastInsertId()
		if err != nil {
			return db.User{}, err
		}

		// update received instead of insert
		if cityID == 0 {
			row, err := q.FindCityByName(ctx, *city)
			if err != nil {
				return db.User{}, err
			}

			cityID = int64(row.ID)
		}

		err = q.InsertUserCity(ctx, db.InsertUserCityParams{UserID: result.ID, CityID: int32(cityID)})
		if err != nil {
			return db.User{}, err
		}
	}

	if err = tx.Commit(); err != nil {
		return db.User{}, err
	}

	return result, nil
}

func (r *Repository) Login(ctx context.Context, user db.User) (db.User, error) {
	row, err := r.queries.FindUserByID(ctx, user.ID)
	if err != nil {
		return db.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(row.Password))
	if err != nil {
		return db.User{}, err
	}

	return db.User{
		ID:        row.ID,
		FirstName: row.FirstName,
		LastName:  row.LastName,
		Gender:    row.Gender,
		Birthday:  row.Birthday,
	}, nil
}
