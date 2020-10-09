package dkr_test

import (
	"bytes"
	"github.com/hekonsek/dkr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsTtyByDefault(t *testing.T) {
	// Given
	out := bytes.NewBufferString("")

	// When
	err := dkr.Sandbox("alpine", nil, []string{"echo", "foo"}, &dkr.SandboxOptions{
		Out: out,
	})

	// Then
	assert.NoError(t, err)
	assert.Equal(t, "foo\n", out.String())
}
