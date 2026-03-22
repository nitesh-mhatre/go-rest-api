package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const SecretKey = "AVSJDSLDMDLydshbdjgdjjdhgkshsuesuwpowpMDMJOUHGSOILJSAJAIJPAJPJAJPKA"

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SecretKey))
}

func ValidateJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, ok := claims["username"].(string)
		if !ok {
			return "", jwt.ErrInvalidKey
		}
		return username, nil
	} else {
		return "", jwt.ErrInvalidKey
	}
}