package dkr_test

import (
	"github.com/hekonsek/dkr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSave(t *testing.T) {
	// Given
	home, err := dkr.NewDkrHome()
	assert.NoError(t, err)
	config := dkr.NewConfig("foo", "bar", []string{"baz"})

	// When
	err = config.Save(home)
	assert.NoError(t, err)

	// Then
	config, err = dkr.ParseConfig(home, "foo")
	assert.NoError(t, err)
	assert.Equal(t, "foo", config.Name)
	assert.Equal(t, "bar", config.Image)
	assert.Equal(t, []string{"baz"}, config.Entrypoint)
}

func TestSaveWithoutPermissionFail(t *testing.T) {
	// Given
	home, err := dkr.NewDkrHomeWihRoot("/")
	assert.NoError(t, err)
	config := dkr.NewConfig("foo", "bar", []string{"baz"})

	// When
	err = config.Save(home)

	// Then
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "denied")
}
