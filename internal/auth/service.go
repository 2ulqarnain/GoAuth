package auth

import (
	"GoAuth/internal/db"
	"context"
	"errors"
)

type AuthService struct {
	repo *Repository
}

func NewAuthService(r *Repository) *AuthService {
	return &AuthService{repo: r}
}

func (s *AuthService) Signup(ctx context.Context, user registerUserPayload) (*db.User, error) {
	hash, err := HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	return s.repo.createUser(ctx, db.CreateUserParams{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: hash,
	})
}

func (s *AuthService) Login(ctx context.Context, user loginPayload) error {
	userByEmail, err := s.repo.getUserByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	isVerified, err := VerifyPassword(hashedPassword, userByEmail.PasswordHash)
	if err != nil {
		return err
	} else if !isVerified {
		return errors.New("incorrect password")
	}
	return nil
}
