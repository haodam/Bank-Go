package token

import (
	"testing"
	"time"

	"github.com/haodam/Bank-Go/util"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	// Create a JWTMaker with a random secret key
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)
	require.NotNil(t, maker, "JWTMaker is nil")

	// Generate random username and token duration
	username := util.RandomOwner()
	duration := time.Minute

	// Record the time when the token is issued
	issuedAt := time.Now()
	// Calculate the expected expiration time
	expiredAt := issuedAt.Add(duration)

	// Create a token
	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	// Verify the token
	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.NotNil(t, payload, "Token verification failed, payload is nil")

	// Perform assertions
	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}
