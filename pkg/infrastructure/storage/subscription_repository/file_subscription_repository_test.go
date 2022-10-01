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

func TestThatAddSubscriptionAndIsSavedWorks(t *testing.T) {
	subsRepo := NewSubscriptionFileRepository()
	emailString := "plainemail@gmail.com"
	base := "BTC"
	quote := "UAH"
	email := *models.NewEmailAddress(emailString)
	pair := *models.NewCurrencyPair(base, quote)
	subscription := *models.NewSubscription(email, pair)
	platform.EnsureFileExists(config.Get().StorageFile)

	subsRepo.AddSubscription(subscription)

	assert.Equal(t, subsRepo.IsSaved(subscription), true)

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
