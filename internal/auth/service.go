package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	secret     string
	expiration time.Duration
}

func NewAuthService() *AuthService {
	exp := time.Duration(3600) * time.Second
	if envExp := os.Getenv("JWT_EXPIRES_IN"); envExp != "" {
		if val, err := time.ParseDuration(envExp + "s"); err == nil {
			exp = val
		}
	}
	return &AuthService{
		secret:     os.Getenv("JWT_SECRET"),
		expiration: exp,
	}
}

func (a *AuthService) GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(a.expiration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.secret))
}

func (a *AuthService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de assinatura inválido")
		}
		return []byte(a.secret), nil
	})
	return token, err
}
