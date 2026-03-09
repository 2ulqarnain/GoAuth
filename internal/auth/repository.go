package auth

import (
	"GoAuth/internal/db"
	"context"
)

type Repository struct {
	db *db.Queries
}

func NewAuthRepository(db *db.Queries) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) createUser(ctx context.Context, user db.CreateUserParams) (*db.User, error) {
	createdUser, err := r.db.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &createdUser, nil
}

func (r *Repository) getUserByEmail(ctx context.Context, email string) (*db.User, error) {
	user, err := r.db.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &db.User{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}, nil
}
