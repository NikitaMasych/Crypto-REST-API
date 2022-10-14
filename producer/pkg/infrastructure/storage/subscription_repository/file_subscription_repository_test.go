package storage

import (
	"os"
	"producer/config"
	"producer/pkg/domain/models"
	"producer/pkg/infrastructure/logger"
	"producer/pkg/utils"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThatAddSubscriptionAndIsSavedWorks(t *testing.T) {
	logger := logger.CreateLogger(logger.TxtLoggerType)
	subsRepo := NewSubscriptionFileRepository(logger)
	emailString := "plainemail@gmail.com"
	base := "BTC"
	quote := "UAH"
	email := *models.NewEmailAddress(emailString)
	pair := *models.NewCurrencyPair(base, quote)
	subscription := *models.NewSubscription(email, pair)
	utils.EnsureFileExists(config.StorageFile)

	subsRepo.AddSubscription(subscription)

	assert.Equal(t, subsRepo.IsSaved(subscription), true)

	cleanup(t)
}

func cleanup(t *testing.T) {
	path := config.StorageFile
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
