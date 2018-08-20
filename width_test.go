package hre

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasDuplication(t *testing.T) {
	assert.True(t, hasDuplication(`a*`))
	assert.True(t, hasDuplication(`b+`))
	assert.True(t, hasDuplication(`c?`))
	assert.False(t, hasDuplication(`d|e`))
}
