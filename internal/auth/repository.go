package auth

import "github.com/jackc/pgx/v5/pgxpool"

type authRepository struct {
	db *pgxpool.Pool
}

func newAuthRepository(db *pgxpool.Pool) *authRepository {
	return &authRepository{
		db: db,
	}
}

//func (r *pgxpool.Pool) createUser(user RegisterPayload) {
//
//}
