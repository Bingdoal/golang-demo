package jwt_service

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const jwtKey = "demo-secret-key"
const jwtTimeoutMinute = 24 * 60

type authClaims struct {
	jwt.StandardClaims
	Body map[string]string
}

func GenerateToken(subject string, body map[string]string) string {
	expiresAt := time.Now().Add(jwtTimeoutMinute * time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   subject,
			ExpiresAt: expiresAt,
		},
		Body: body,
	})
	tokenString, _ := token.SignedString([]byte(jwtKey))
	return tokenString
}

func ValidateToken(tokenString string) (string, map[string]string, error) {
	var claims authClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		return "", map[string]string{}, err
	}
	if !token.Valid {
		return "", map[string]string{}, errors.New("invalid token")
	}
	body := claims.Body

	return claims.Subject, body, nil
}
