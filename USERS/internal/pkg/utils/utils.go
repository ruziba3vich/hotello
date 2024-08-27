package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"github.com/ruziba3vich/hotello-users/internal/pkg/config"
)

type (
	TokenGenerator struct {
		secretKey string
	}
	PasswordHasher struct{}
)

func NewTokenGenerator(cfg *config.Config) *TokenGenerator {
	return &TokenGenerator{
		secretKey: cfg.GetSecretKey(),
	}
}

func (t *TokenGenerator) GenerateToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(t.secretKey))
	if err != nil {
		return "", fmt.Errorf("could not create token: %s", err.Error())
	}

	return tokenString, nil
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

func (p *PasswordHasher) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("could not hash password: %s", err.Error())
	}
	return string(hash), nil
}

func (p *PasswordHasher) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewPasswordHasher() *PasswordHasher {
	return &PasswordHasher{}
}
