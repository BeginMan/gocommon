package random

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestRandom(t *testing.T)  {
	assert.Len(t, String(32), 32)
	r := New()
	assert.Regexp(t, regexp.MustCompile("[0-9]+$"), r.String(8, Numeric))
}
