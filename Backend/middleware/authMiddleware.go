package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// ambil token
	tokenString, err := c.Cookie("jwt")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized"})
		return
	}

	//
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	//
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userID := uint(claims["user_id"].(float64))

		c.Set("user_id", userID)

		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "fail read token"})
	}

}
