package middleware

import (
	"log"
	"net/http"
	"strings"

	"Go-Auth-API/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	cfg, err := config.LoadJWTConfig()
	if err != nil {
		panic("JWT configuration error: " + err.Error())
	}
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is required",
			})
			return
		}

		parts := strings.Fields(authHeader)
		if len(parts) != 2 ||
			!strings.EqualFold(parts[0], "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization header format",
			})
			return
		}

		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(
			parts[1],
			claims,
			func(token *jwt.Token) (interface{}, error) {
				return cfg.Secret, nil
			},
			jwt.WithValidMethods([]string{
				jwt.SigningMethodHS256.Alg(),
			}),
			jwt.WithExpirationRequired(),
		)

		// if err != nil || token == nil || !token.Valid {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		// 		"error": "Invalid or expired token",
		// 	})
		// 	return
		// }

		if err != nil {
	log.Printf("JWT validation failed: %v", err)

	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": "Invalid or expired token",
	})
	return
}

if token == nil || !token.Valid {
	log.Println("JWT validation failed: token is nil or invalid")

	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": "Invalid or expired token",
	})
	return
}

		id, ok := claims["id"].(float64)
		if !ok || id <= 0 || id != float64(uint(id)) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid user ID in token",
			})
			return
		}

		c.Set("userID", uint(id))
		c.Next()
	}
}