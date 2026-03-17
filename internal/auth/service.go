package auth

import (
	"GoAuth/internal/db"
	"GoAuth/internal/errs"
	"context"
	"fmt"
	"time"
)

type Service struct {
	repo *Repository
	jwt  *JWTManager
}

func NewAuthService(r *Repository, jwt *JWTManager) *Service {
	return &Service{repo: r, jwt: jwt}
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

func (s *Service) Login(ctx context.Context, user loginPayload) (string, string, error) {
	userByEmail, err := s.repo.getUserByEmail(ctx, user.Email)
	if err != nil {
		return "", "", fmt.Errorf("svc Login -> %v", err)
	}
	isVerified, err := VerifyPassword(userByEmail.PasswordHash, user.Password)
	if err != nil {
		return "", "", fmt.Errorf("svc Login VerifyPass: %v", err)
	} else if !isVerified {
		return "", "", errs.ErrInvalidPassword
	}

	accessToken, _ := s.jwt.GenerateToken(userByEmail.ID, time.Minute*15)
	_, refreshTokenHash, err := GenerateRefreshToken()

	return accessToken, refreshTokenHash, nil
}

func (s *Service) RenewAccessToken(ctx context.Context, refreshToken string) (string, error) {
	tokenHash := HashRefreshToken(refreshToken)

}
