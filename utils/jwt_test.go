package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenAndParseToken(t *testing.T) {
	secret := "secret"
	issuer := "issuer"
	expire := 10
	token, err := GenToken(secret, issuer, expire, 1, "name")
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	mc, err := ParseToken(secret, token)
	assert.Nil(t, err)
	assert.Equal(t, 1, mc.Id)
}

func TestParseTokenError(t *testing.T) {
	tokenString := "tokenString"

	_, err := ParseToken("secret", tokenString)

	assert.ErrorIs(t, err, ErrTokenMalformed)
}
