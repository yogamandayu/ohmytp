package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Secret        string
	SigningMethod jwt.SigningMethod
}

// NewJWT is a constructor.
func NewJWT(secret string) *JWT {
	return &JWT{
		Secret:        secret,
		SigningMethod: jwt.SigningMethodHS256,
	}
}

// Generate creates a new JWT token with the given payload.
func (j *JWT) Generate(payload map[string]interface{}) (string, error) {
	mapClaims := jwt.MapClaims{}
	for key, val := range payload {
		mapClaims[key] = val
	}

	token := jwt.NewWithClaims(j.SigningMethod, mapClaims)

	tokenString, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateToken parses and validates a JWT token string.
func (j *JWT) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("jwt.error.unexpected_signing_method")
		}
		return []byte(j.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("jwt.error.invalid_token")
	}

	return claims, nil
}
