package repository

import (
	"GenesisTask/config"
	"GenesisTask/model"
	"GenesisTask/platform"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThatAddAndIsExistWorks(t *testing.T) {
	userRepo := New()
	email := "plainemail@gmail.com"
	user := model.NewUser(email)
	platform.EnsureFileExists(config.Get().StorageFile)

	userRepo.Add(user)

	assert.Equal(t, userRepo.IsExist(user), true)

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
