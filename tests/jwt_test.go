package tests

import (
	"testing"
	"time"

	"github.com/saddam-satria/go-boilerplate/pkg"
	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {
	id := "123"
	email := "test@gmail.com"
	expired := time.Now().Add(time.Minute * 60 * 24).Unix()
	token, err := pkg.GenerateToken(id, email, expired, "inisecretkey")

	assert.NoError(t, err, "error generate token")

	claims, err := pkg.VerifyToken(token)

	assert.NoError(t, err, "error verifiy token")

	assert.Equal(t, id, claims["id"], token)
}
