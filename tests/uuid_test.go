package tests

import (
	"testing"

	"github.com/saddam-satria/go-boilerplate/pkg"
	"github.com/stretchr/testify/assert"
)

func TestUUID(t *testing.T) {

	generatedUUID := pkg.GenerateUUID()

	isValid := pkg.IsUUID(generatedUUID)

	assert.Equal(t, true, isValid)
}
