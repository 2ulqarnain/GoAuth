package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

// GenerateRefreshToken generates string and hash of refresh token.
// first value in returned values is string and the second one is hashed string
func GenerateRefreshToken() (string, string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", "", err
	}

	tokenString := base64.RawURLEncoding.EncodeToString(token)

	hash := sha256.Sum256([]byte(tokenString))
	hashString := base64.RawURLEncoding.EncodeToString(hash[:])

	return tokenString, hashString, nil
}

// HashRefreshToken hashes given token
func HashRefreshToken(tokenString string) string {
	hash := sha256.Sum256([]byte(tokenString))

	return base64.RawURLEncoding.EncodeToString(hash[:])
}
