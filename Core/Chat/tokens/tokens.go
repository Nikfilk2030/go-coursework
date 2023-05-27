package tokens

import (
	"errors"
	"fmt"
	"time"

	jwt4 "github.com/golang-jwt/jwt/v4"
)

var ErrInvalidToken = errors.New("invalid token")

func CreateToken(subject string, expirationTime time.Time, secret string) (string, error) {
	claims := jwt4.RegisteredClaims{
		Issuer:    "auth",
		ExpiresAt: jwt4.NewNumericDate(expirationTime),
		Subject:   subject,
	}
	token := jwt4.NewWithClaims(jwt4.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("could not sign token: %w", err)
	}
	return tokenString, nil
}

func ParseToken(tokenString string, secret string) (*jwt4.Token, error) {
	keyFunc := func(t *jwt4.Token) (interface{}, error) {
		if t.Method.Alg() != jwt4.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}

		return []byte(secret), nil
	}
	token, err := jwt4.Parse(tokenString, keyFunc)
	if err != nil {
		if errors.Is(err, jwt4.ErrTokenExpired) {
			return token, fmt.Errorf("token is expired: %w", err)
		}
		return nil, fmt.Errorf("could not parse token: %w", err)
	}
	if !token.Valid {
		return nil, ErrInvalidToken
	}
	return token, nil
}
