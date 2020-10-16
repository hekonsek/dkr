package dkr_test

import (
	dkr2 "github.com/hekonsek/dkr/dkr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSave(t *testing.T) {
	// Given
	home, err := dkr2.NewDkrHome()
	assert.NoError(t, err)
	config := dkr2.NewConfig("foo", "bar", []string{"baz"})

	// When
	err = config.Save(home)
	assert.NoError(t, err)

	// Then
	config, err = dkr2.ParseConfig(home, "foo")
	assert.NoError(t, err)
	assert.Equal(t, "foo", config.Name)
	assert.Equal(t, "bar", config.Image)
	assert.Equal(t, []string{"baz"}, config.Entrypoint)
}

func TestParseNotExistingCommand(t *testing.T) {
	// Given
	home, err := dkr2.NewDkrHome()
	assert.NoError(t, err)

	// When
	config, err := dkr2.ParseConfig(home, "noSuchCommand")

	// Then
	assert.NoError(t, err)
	assert.Nil(t, config)
}

func TestSaveWithoutPermissionFail(t *testing.T) {
	// Given
	home, err := dkr2.NewDkrHomeWihRoot("/")
	assert.NoError(t, err)
	config := dkr2.NewConfig("foo", "bar", []string{"baz"})

	// When
	err = config.Save(home)

	// Then
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "denied")
}

func TestImportConfigYml(t *testing.T) {
	// When
	configYml, err := dkr2.ImportConfigYml("terraform")
	assert.NoError(t, err)

	// Then
	assert.Equal(t, "image: hekonsek/dkr-terraform", string(configYml))
}
