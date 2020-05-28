package dkr_test

import (
	"github.com/hekonsek/dkr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateHome(t *testing.T) {
	home, err := dkr.NewDkrHome()
	assert.NoError(t, err)
	assert.DirExists(t, home.Root)
}