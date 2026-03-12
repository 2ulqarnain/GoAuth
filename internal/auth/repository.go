package auth

import (
	"GoAuth/internal/db"
	"context"
	"fmt"
)

type Repository struct {
	db *db.Queries
}

func NewAuthRepository(db *db.Queries) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) createUser(ctx context.Context, user db.CreateUserParams) (*db.CreateUserRow, error) {
	createdUser, err := r.db.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("repo createUser: %v", err)
	}
	return &createdUser, nil
}

func (r *Repository) getUserByEmail(ctx context.Context, email string) (*db.User, error) {
	user, err := r.db.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("repo getUserByEmail: %v", err)
	}
	return &db.User{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}, nil
}
