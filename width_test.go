package hre

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcMaxWidth(t *testing.T) {
	assert.Equal(t, [2]int{1, 1}, calcWidthRange(`a`))
	assert.Equal(t, [2]int{4, 4}, calcWidthRange(`a...`))
	assert.Equal(t, [2]int{1, 1}, calcWidthRange(`a|b`))
	assert.Equal(t, [2]int{1, 1}, calcWidthRange(`((a)|b)`))
	assert.Equal(t, [2]int{0, 1}, calcWidthRange(`((a)|b)?`))
	assert.Equal(t, [2]int{1, 1}, calcWidthRange(`[abcde]`))
	assert.Equal(t, [2]int{1, 1}, calcWidthRange(`[\]abcde]`))
	assert.Equal(t, [2]int{1, 1}, calcWidthRange(`[(abcde)]`))
	assert.Equal(t, [2]int{3, 126}, calcWidthRange(`(...){1,42}`))
	assert.Equal(t, [2]int{0, -1}, calcWidthRange(`.*`))
	assert.Equal(t, [2]int{1, -1}, calcWidthRange(`.+`))
	assert.Equal(t, [2]int{0, -1}, calcWidthRange(`.*?`))
	assert.Equal(t, [2]int{1, -1}, calcWidthRange(`.+?`))
}
