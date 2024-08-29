package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ruziba3vich/hotello-gateway/internal/pkg/config"
)

type (
	TokenGenerator struct {
		secretKey string
	}
)

func New(cfg *config.Config) *TokenGenerator {
	return &TokenGenerator{
		secretKey: cfg.GetSecretKey(),
	}
}

func (t *TokenGenerator) VerifyToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(t.secretKey), nil
	})

	if err != nil {
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			return time.Now().Unix() <= int64(exp)
		}
	}

	return false
}
