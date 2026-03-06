package tokens

import "time"

type RefreshToken struct {
	Id        int32     `json:"id,required"`
	UserId    int       `json:"user_id,required"`
	TokenHash string    `json:"token_hash,required"`
	CreatedAt time.Time `json:"created_at,required"`
	ExpiresAt time.Time `json:"expires_at,required"`
}
