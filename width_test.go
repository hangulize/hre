package hre

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasRepeatition(t *testing.T) {
	assert.True(t, hasRepeatition(`a*`))
	assert.True(t, hasRepeatition(`b+`))
	assert.True(t, hasRepeatition(`c?`))
	assert.False(t, hasRepeatition(`d|e`))
	assert.False(t, hasRepeatition(`...`))
}
