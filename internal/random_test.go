package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomString(t *testing.T) {
	for i := 0; i < 1000; i++ {
		str := RandomString(16)
		t.Logf("%s\n", RandomString(16))
		assert.True(t, len(str) == 16, "len error")
	}
}
