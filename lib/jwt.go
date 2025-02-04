package lib

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateAccessToken(id uint, role string, isVerrified bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat":      jwt.NewNumericDate(time.Now()),
		"id":       id,
		"role":     role,
		"verified": isVerrified,
		"exp":      jwt.NewNumericDate(time.Now().Add(time.Second * 30)), // 30 seconds
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateRefreshToken(id uint, role string, isVerrified bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat":      jwt.NewNumericDate(time.Now()),
		"id":       id,
		"role":     role,
		"verified": isVerrified,
		"exp":      jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 24 hours
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func GenerateTokens(id uint, role string, isVerrified bool) (string, string, error) {
	accessToken, err := GenerateAccessToken(id, role, isVerrified)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := GenerateRefreshToken(id, role, isVerrified)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
