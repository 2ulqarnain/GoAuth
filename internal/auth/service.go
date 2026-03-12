package auth

import (
	"GoAuth/internal/db"
	"GoAuth/internal/errs"
	"context"
	"fmt"
)

type Service struct {
	repo *Repository
}

func NewAuthService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Signup(ctx context.Context, user signupPayload) (*db.CreateUserRow, error) {
	hash, err := HashPassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("svc Signup: %v", err)
	}

	return s.repo.createUser(ctx, db.CreateUserParams{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: hash,
	})
}

func (s *Service) Login(ctx context.Context, user loginPayload) error {
	userByEmail, err := s.repo.getUserByEmail(ctx, user.Email)
	if err != nil {
		return fmt.Errorf("svc Login -> %v", err)
	}
	isVerified, err := VerifyPassword(userByEmail.PasswordHash, user.Password)
	if err != nil {
		return fmt.Errorf("svc Login VerifyPass: %v", err)
	} else if !isVerified {
		return errs.ErrInvalidPassword
	}
	return nil
}
