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

	createUserParams := db.CreateUserParams{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  string(pass),
		Birthday:  user.Birthday,
	}
	if biography != nil {
		createUserParams.Biography = sql.NullString{String: *biography, Valid: true}
	}

	sqlResult, err := r.queries.CreateUser(ctx, createUserParams)
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

	return result, nil
}

func (r *Repository) Login(ctx context.Context, user db.User) (db.User, error) {
	row, err := r.queries.FindUserByID(ctx, user.ID)
	if err != nil {
		return db.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(row.Password), []byte(user.Password))
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
