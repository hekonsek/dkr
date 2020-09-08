package dkr_test

import (
	"github.com/hekonsek/dkr/dkr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadingBashrc(t *testing.T) {
	// When
	bashrc, err := dkr.NewBashrc()
	assert.NoError(t, err)

	// Then
	assert.NotEmpty(t, bashrc.Lines())
}

func TestNotHaveAlias(t *testing.T) {
	// Given
	bashrc, err := dkr.NewBashrc()
	assert.NoError(t, err)

	// When
	hasAlias, err := bashrc.HasAlias("someRandomAlias")
	assert.NoError(t, err)

	// Then
	assert.False(t, hasAlias)
}

func TestAddAlias(t *testing.T) {
	// Given
	command := "foo"
	bashrc, err := dkr.NewBashrc()
	assert.NoError(t, err)

	// When
	err = bashrc.AddAlias(command)
	assert.NoError(t, err)

	// Then
	hasAlias, err := bashrc.HasAlias(command)
	assert.NoError(t, err)
	assert.True(t, hasAlias)
}
