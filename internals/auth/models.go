package auth

import "time"

type LoginPayload struct {
	Username string `json:"username,required"`
	Password string `json:"password,required"`
}

type User struct {
	Id           int       `json:"id,required"`
	Name         string    `json:"name,required"`
	Email        string    `json:"email,required"`
	PasswordHash string    `json:"password_hash,required"`
	CreatedAt    time.Time `json:"created_at,required"`
}
