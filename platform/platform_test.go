package platform

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsureFileExists(t *testing.T) {
	path := "somedir/somefile.txt"

	EnsureFileExists(path)

	assert.FileExists(t, path)

	os.RemoveAll("somedir")
}

func TestEnsureDirectoryExists(t *testing.T) {
	directory := "somedir"

	ensureDirectoryExists(directory)

	assert.DirExists(t, directory)

	os.RemoveAll(directory)
}
