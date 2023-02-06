package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"go01-airbnb/config"
	"time"
)

type Token struct {
	AccessToken string `json:"access_token"`
}

type TokenPayload struct {
	Email string `json:"email"`
}

type myClaim struct {
	Payload TokenPayload `json:"payload"`
	jwt.RegisteredClaims
}

func GenerateJWT(data TokenPayload, secretKey string) (*Token, error) {
	expiresAt := jwt.NewNumericDate(time.Now().Add(time.Hour * 12))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim{
		data,
		jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	accessToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken: accessToken,
	}, nil
}

func ValidateJWT(accessToken string, cfg *config.Config) (*TokenPayload, error) {
	token, err := jwt.ParseWithClaims(accessToken, &myClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.App.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claim, ok := token.Claims.(*myClaim)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return &claim.Payload, nil
}
