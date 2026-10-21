package services

import (
	"fmt"
	"time"

	"Go-Auth-API/config"

	"github.com/golang-jwt/jwt/v5"
)

type JWTServiceInterface interface {
	GenerateAccessToken(userID uint) (string, error)
}

type AccessClaims struct {
	UserID uint `json:"id"`
	jwt.RegisteredClaims
}

type JWTService struct {
	config *config.JWTConfig
}

var _ JWTServiceInterface = (*JWTService)(nil)

func NewJWTService(cfg *config.JWTConfig) *JWTService {
	return &JWTService{
		config: cfg,
	}
}

func (jwtS *JWTService) GenerateAccessToken(userID uint) (string, error) {
	now := time.Now()

	claims := AccessClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(jwtS.config.AccessTokenTTL)),
			Subject:   fmt.Sprintf("%d", userID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtS.config.Secret)
	if err != nil {
		return "", fmt.Errorf("sign access token: %w", err)
	}

	return signedToken, nil
}