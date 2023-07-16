package users

import (
	"context"
	"github.com/magmel48/social-network/internal/db"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func (r *Repository) Register(ctx context.Context, user db.User, city, hobbies *string) (db.User, error) {
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

	sqlResult, err := q.CreateUser(ctx, db.CreateUserParams{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  string(pass),
		Birthday:  user.Birthday,
	})
	if err != nil {
		return db.User{}, err
	}

	id, err := sqlResult.LastInsertId()
	if err != nil {
		return db.User{}, err
	}

	result := db.User{
		ID:        int32(id),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Gender:    user.Gender,
		Birthday:  user.Birthday,
	}

	if city != nil {
		err := q.UpsertCity(ctx, *city)
		if err != nil {
			return db.User{}, err
		}
	}

	if hobbies != nil {
		hbs := strings.Split(*hobbies, " ")
		for _, h := range hbs {
			err := q.UpsertHobby(ctx, h)
			if err != nil {
				return db.User{}, err
			}
		}
	}

	if err = tx.Commit(); err != nil {
		return db.User{}, err
	}

	return result, nil
}

func (r *Repository) Login(ctx context.Context, user db.User) (db.User, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return db.User{}, err
	}

	row, err := r.queries.FindUserWithCheckingPassword(ctx, db.FindUserWithCheckingPasswordParams{
		ID:       user.ID,
		Password: string(p),
	})
	if err != nil {
		return db.User{}, err
	}

	result := db.User{
		ID:        user.ID,
		FirstName: row.FirstName,
		LastName:  row.LastName,
		Gender:    row.Gender,
		Birthday:  row.Birthday,
		CreatedAt: row.CreatedAt,
	}

	return result, nil
}
