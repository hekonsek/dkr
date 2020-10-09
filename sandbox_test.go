package dkr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsTtyByDefault(t *testing.T) {
	tty := tty()
	assert.True(t, tty)
}
