package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	userID := "test-user-id"

	// Generate a JWT token for testing
	token, err := GenerateToken(userID)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Print the token
	t.Log("Generated Token:", token)
}
