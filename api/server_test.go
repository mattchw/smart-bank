package api

import (
	"testing"
	"time"

	db "github.com/mattchw/smart-bank/db/sqlc"
	"github.com/mattchw/smart-bank/util"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}
