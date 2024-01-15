package tests

import (
	"testing"

	"github.com/saddam-satria/go-boilerplate/pkg"
	"github.com/stretchr/testify/assert"
)

func TestBcrypt(t *testing.T) {
	password := "test123"

	hashedPassword, err := pkg.HashedBcrypt(password)

	assert.NoError(t, err)

	isMatch := pkg.VerifyHashed(password, hashedPassword)

	assert.Equal(t, true, isMatch)
}
