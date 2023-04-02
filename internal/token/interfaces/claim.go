package token_interfaces

import (
	"time"

	"github.com/mattchw/smart-bank/errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claim struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
	jwt.RegisteredClaims
}

func NewClaim(username string, duration time.Duration) (*Claim, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	claim := &Claim{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return claim, nil
}

func (c *Claim) Validate() error {
	if time.Now().After(c.ExpiredAt) {
		return errors.ErrExpiredToken
	}
	return nil
}
