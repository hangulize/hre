package hre

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcMaxWidth(t *testing.T) {
	assert.Equal(t, 1, calcMaxWidth(`a`))
	assert.Equal(t, 4, calcMaxWidth(`a...`))
	assert.Equal(t, 1, calcMaxWidth(`a|b`))
	assert.Equal(t, 1, calcMaxWidth(`((a)|b)`))
	assert.Equal(t, 1, calcMaxWidth(`((a)|b)?`))
	assert.Equal(t, 3, calcMaxWidth(`a|bc|def|g|hi|jkl`))
	assert.Equal(t, 1, calcMaxWidth(`[abcde]`))
	assert.Equal(t, 1, calcMaxWidth(`[\]abcde]`))
	assert.Equal(t, 1, calcMaxWidth(`[(abcde)]`))
	assert.Equal(t, 126, calcMaxWidth(`(...){1,42}`))
	assert.Equal(t, -1, calcMaxWidth(`.*`))
	assert.Equal(t, -1, calcMaxWidth(`.+`))
	assert.Equal(t, -1, calcMaxWidth(`.*?`))
	assert.Equal(t, -1, calcMaxWidth(`.+?`))
	assert.Equal(t, -1, calcMaxWidth(`(.+|...)`))
	assert.Equal(t, -1, calcMaxWidth(`.+...`))
	assert.Equal(t, 0, calcMaxWidth(`??INVALID??REGEXP??`))
}
