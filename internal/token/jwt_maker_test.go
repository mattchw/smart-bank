package token

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mattchw/smart-bank/errors"
	token_interfaces "github.com/mattchw/smart-bank/internal/token/interfaces"
	"github.com/mattchw/smart-bank/util"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomName()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestJWTMaker_InvalidSecret(t *testing.T) {
	_, err := NewJWTMaker(util.RandomString(31))
	require.Error(t, err)
}

func TestJWTMaker_ExpiredToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomName()
	duration := -time.Minute

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, fmt.Sprintf("token has invalid claims: token is expired, %s", errors.ErrExpiredToken.Error()))
	require.Nil(t, payload)
}

func TestJWTMaker_InvalidToken_NoneAlg(t *testing.T) {
	payload, err := token_interfaces.NewClaim(util.RandomName(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, fmt.Sprintf("token is unverifiable: error while executing keyfunc: %s", errors.ErrInvalidToken.Error()))
	require.Nil(t, payload)
}
