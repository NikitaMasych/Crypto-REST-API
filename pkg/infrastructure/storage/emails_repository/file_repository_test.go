package storage

import (
	"GenesisTask/config"
	"GenesisTask/pkg/domain/models"
	"GenesisTask/pkg/platform"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThatAddEmailAndIsSavedWorks(t *testing.T) {
	userRepo := NewFileRepository()
	email := "plainemail@gmail.com"
	address := models.NewEmailAddress(email)
	platform.EnsureFileExists(config.Get().StorageFile)

	userRepo.AddEmail(*address)

	assert.Equal(t, userRepo.IsSaved(*address), true)

	cleanup(t)
}

func cleanup(t *testing.T) {
	path := config.Get().StorageFile
	_, err := os.Stat(path)
	if err != nil {
		t.Error(err)
	}
	directory := strings.Split(path, "/")[0]
	err = os.RemoveAll(directory)
	if err != nil {
		t.Error(err)
	}
}
