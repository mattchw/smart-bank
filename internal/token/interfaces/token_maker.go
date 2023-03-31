package token_interfaces

import (
	"time"
)

type TokenMaker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Claim, error)
}
