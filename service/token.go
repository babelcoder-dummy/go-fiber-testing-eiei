package service

import (
	"fmt"
	"time"

	"github.com/babelcoder-enterprise-courses/go-fiber-testing/config"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/lo"
)

type Token struct {
}

func (t *Token) Create(payload map[string]any, expiresIn, secretKey string) (string, error) {
	duration, err := time.ParseDuration(expiresIn)
	if err != nil {
		return "", err
	}

	expiresTime := time.Now().Add(duration)
	claims := jwt.MapClaims(lo.Assign(
		map[string]any{
			"exp": jwt.NewNumericDate(expiresTime),
		},
		payload,
	))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return ss, nil
}

func (t *Token) VerifyToken(tokenString, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func (t *Token) GenerateAccessAndRefreshTokens(uid uint, role model.Role) (string, string, error) {
	userID := uid
	accessTokenExpiresIn := config.Env.Expiration.AccessToken
	accessTokenSecretKey := config.Env.Secret.AccessToken
	claims := map[string]any{"sub": userID, "role": role}
	accessToken, err := t.Create(claims, accessTokenExpiresIn, accessTokenSecretKey)
	if err != nil {
		return "", "", fmt.Errorf("cannot create access token: %w", err)
	}

	refreshTokenExpiresIn := config.Env.Expiration.RefreshToken
	refreshTokenSecretKey := config.Env.Secret.RefreshToken
	refreshToken, err := t.Create(claims, refreshTokenExpiresIn, refreshTokenSecretKey)
	if err != nil {
		return "", "", fmt.Errorf("cannot create refresh token: %w", err)
	}

	return accessToken, refreshToken, nil
}
