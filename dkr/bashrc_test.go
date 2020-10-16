package dkr_test

import (
	"github.com/hekonsek/dkr/dkr"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path"
	"testing"
)

func TestLoadingBashrc(t *testing.T) {
	// When
	tmpBashrc, err := ioutil.TempFile("", "dkr-test-*")
	assert.NoError(t, err)
	bashrc, err := dkr.NewBashrcFromFile(tmpBashrc.Name())
	assert.NoError(t, err)

	// Then
	assert.NotEmpty(t, bashrc.Lines())
}

func TestNotHavePath(t *testing.T) {
	// Given
	tmpBashrc, err := ioutil.TempFile("", "dkr-test-*")
	assert.NoError(t, err)
	bashrc, err := dkr.NewBashrcFromFile(tmpBashrc.Name())
	assert.NoError(t, err)

	// When
	hasAlias := bashrc.HasPath()
	assert.NoError(t, err)

	// Then
	assert.False(t, hasAlias)
}

func TestHasPath(t *testing.T) {
	// Given
	tmpBashrc, err := ioutil.TempFile("", "dkr-test-*")
	assert.NoError(t, err)
	bashrc, err := dkr.NewBashrcFromFile(tmpBashrc.Name())
	assert.NoError(t, err)
	err = bashrc.AddPath("/tmp")
	assert.NoError(t, err)

	// When
	hasPath := bashrc.HasPath()

	// Then
	assert.True(t, hasPath)
}

func TestAddCommandProxyToPath(t *testing.T) {
	// Given
	tmpBashrc, err := ioutil.TempFile("", "dkr-test-*")
	assert.NoError(t, err)
	bashrc, err := dkr.NewBashrcFromFile(tmpBashrc.Name())
	assert.NoError(t, err)
	tmpHome, err := ioutil.TempDir("", "dkr-test-*")
	assert.NoError(t, err)
	home, err := dkr.NewDkrHomeWihRoot(tmpHome)
	assert.NoError(t, err)

	// When
	err = bashrc.AddCommandProxy(home.Bin(), "foo")
	assert.NoError(t, err)

	// Then
	proxyPath := path.Join(home.Bin(), "foo")
	assert.FileExists(t, proxyPath)
	proxyBytes, err := ioutil.ReadFile(proxyPath)
	assert.NoError(t, err)
	assert.Contains(t, string(proxyBytes), "#!/bin/bash\n")
	assert.Contains(t, string(proxyBytes), "\ndkr run foo")
}
