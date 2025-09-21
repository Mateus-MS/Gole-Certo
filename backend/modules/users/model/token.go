package user_model

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

type Token struct {
	Token     string    `json:"token"     binding:"required" bson:"token"`
	CreatedAt time.Time `json:"createdAt" binding:"required" bson:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt" binding:"required" bson:"expiresAt"`
}

func (t *Token) IsValid() bool {
	return time.Now().Before(t.ExpiresAt)
}

func NewToken(tokenCost int, duration time.Duration) (Token, error) {
	var token Token

	tokenString, err := GenerateRandomToken(tokenCost)
	if err != nil {
		return token, err
	}

	token.Token = tokenString
	token.CreatedAt = time.Now()
	token.ExpiresAt = time.Now().Add(duration)

	return token, nil
}

func GenerateRandomToken(n int) (string, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}
