package config

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type JWTConfig struct {
	Secret        []byte
	AccessTokenTTL time.Duration
}

func LoadJWTConfig() (*JWTConfig, error) {
	secret := strings.TrimSpace(os.Getenv("JWT_SECRET"))
	if len(secret) < 32 {
		return nil, fmt.Errorf("JWT_SECRET must be at least 32 characters")
	}

	ttl := strings.TrimSpace(os.Getenv("JWT_ACCESS_TOKEN_TTL"))
	if ttl == "" {
		ttl = "15m"
	}

	duration, err := time.ParseDuration(ttl)
	if err != nil || duration <= 0 {
		return nil, fmt.Errorf("invalid JWT_ACCESS_TOKEN_TTL")
	}

	return &JWTConfig{
		Secret:         []byte(secret),
		AccessTokenTTL: duration,
	}, nil
}