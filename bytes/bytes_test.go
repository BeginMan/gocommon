package bytes

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBufferRead(t *testing.T) {
  buf:=bytes.NewBuffer([]byte{'h','e','l','l','o'})
  str, _ := ReadString(buf)
  assert.Equal(t, "", str)
}

