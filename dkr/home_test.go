package dkr_test

import (
	dkr2 "github.com/hekonsek/dkr/dkr"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestCreateHome(t *testing.T) {
	// Given
	tempHome, err := ioutil.TempDir("", "dkr-test-*")
	assert.NoError(t, err)

	// When
	home, err := dkr2.NewDkrHomeWihRoot(tempHome)

	// Then
	assert.NoError(t, err)
	assert.DirExists(t, home.Root())
}

func TestCreateBin(t *testing.T) {
	// Given
	tempHome, err := ioutil.TempDir("", "dkr-test-*")
	assert.NoError(t, err)

	// When
	home, err := dkr2.NewDkrHomeWihRoot(tempHome)

	// Then
	assert.NoError(t, err)
	assert.DirExists(t, home.Bin())
}
