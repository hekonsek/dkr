package dkr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPossibleConfigUrlsContainHekonsekConfigFileInMainBranch(t *testing.T) {
	// When
	urls := generatePossibleConfigUrls("foo")

	// Then
	assert.Contains(t, urls, "https://raw.githubusercontent.com/hekonsek/dkr-foo/main/config.yml")
}

func TestPossibleConfigUrlsContainHekonsekConfigFileInMasterBranch(t *testing.T) {
	// When
	urls := generatePossibleConfigUrls("foo")

	// Then
	assert.Contains(t, urls, "https://raw.githubusercontent.com/hekonsek/dkr-foo/master/config.yml")
}
