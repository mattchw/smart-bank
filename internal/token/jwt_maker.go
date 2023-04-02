package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mattchw/smart-bank/errors"
	token_interfaces "github.com/mattchw/smart-bank/internal/token/interfaces"
)

const minSecretLength = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (token_interfaces.TokenMaker, error) {
	if len(secretKey) < minSecretLength {
		return nil, fmt.Errorf("secret key must be at least %d characters long", minSecretLength)
	}
	return &JWTMaker{secretKey: secretKey}, nil
}

func (j *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := token_interfaces.NewClaim(username, duration)
	if err != nil {
		return "", err
	}

	claims := token_interfaces.Claim{
		payload.ID,
		payload.Username,
		payload.IssuedAt,
		payload.ExpiredAt,
		jwt.RegisteredClaims{
			Issuer:    "smart-bank",
			ExpiresAt: jwt.NewNumericDate(payload.ExpiredAt),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString([]byte(j.secretKey))
}

func (j *JWTMaker) VerifyToken(token string) (*token_interfaces.Claim, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.ErrInvalidToken
		}
		return []byte(j.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &token_interfaces.Claim{}, keyFunc)
	if err != nil {
		return nil, err
	}

	payload, ok := jwtToken.Claims.(*token_interfaces.Claim)
	if !ok && !jwtToken.Valid {
		return nil, errors.ErrInvalidToken
	}

	return payload, nil
}
