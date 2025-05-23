package utils

import (
	"errors"
	"time"

	env "github.com/ghozir/sighapp/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateJWT(userID string) (string, string, error) {
	secret := env.Config.JWTSecret
	if secret == "" {
		return "", "", jwt.ErrTokenMalformed
	}

	jti := uuid.NewString() // 🎯 Generate UUID

	claims := jwt.MapClaims{
		"sub": userID,
		"jti": jti, // ✅ Simpan jti ke dalam token
		"exp": time.Now().Add(env.Config.JWTExpiresIn).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", "", err
	}

	return signedToken, jti, nil
}

func DecodeJWT(tokenString string) (userID string, jti string, err error) {
	secret := env.Config.JWTSecret
	if secret == "" {
		return "", "", jwt.ErrTokenMalformed
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", "", errors.New("invalid token")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return "", "", errors.New("userID not found in token")
	}

	jtiVal, ok := claims["jti"].(string)
	if !ok {
		return "", "", errors.New("jti not found in token")
	}

	return sub, jtiVal, nil
}
