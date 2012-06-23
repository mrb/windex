package windex

import (
	//"github.com/mrb/windex"
	"github.com/bmizerany/assert"
	"testing"
)

var (
	LogLine = []byte(`[2012-06-23T00:00:00+00:00] Boom  Boom  Boom  Boom  Boom  Boom  Boom  Boom  Boom`)
	Lines   = 50
)

func setupInputBuffer(t *testing.T) (input []byte) {
	input = make([]byte, 0)
	assert.T(t, input != nil)
	return
}

func setupOutputBuffer(t *testing.T) (output []byte) {
	output = make([]byte, 0)
	assert.T(t, output != nil)
	return
}

func TestFillInputBuffer(t *testing.T) {
	input := setupInputBuffer(t)
	assert.T(t, input != nil)

	for i := 0; i < Lines; i++ {
		input = append(input, LogLine...)
	}

	assert.T(t, len(input) == (Lines*len(LogLine)))
}
