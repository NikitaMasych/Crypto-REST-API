package storage

import (
	"GenesisTask/config"
	"GenesisTask/pkg/application"
	"GenesisTask/pkg/domain/models"
	"GenesisTask/pkg/errors"
	"bufio"
	"log"
	"os"
	"strings"
)

type SubscriptionFileRepository struct {
	path string
}

func NewSubscriptionFileRepository() application.SubscriptionStorage {
	path := config.StorageFile
	return &SubscriptionFileRepository{path}
}

func (r *SubscriptionFileRepository) IsSaved(subscription models.Subscription) bool {
	file, err := os.OpenFile(r.path, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if subscription.ToString() == scanner.Text() {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return false
}

func (r *SubscriptionFileRepository) AddSubscription(subscription models.Subscription) error {
	file, err := os.OpenFile(r.path, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(subscription.ToString() + "\n")
	if err != nil {
		return err
	}
	err = file.Sync()
	if err != nil {
		return err
	}

	return err
}

func (r *SubscriptionFileRepository) GetAll() []models.Subscription {
	file, err := os.OpenFile(r.path, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var subscriptions []models.Subscription
	for scanner.Scan() {
		line := scanner.Text()
		subscription, err := ComposeSubscription(line)
		if err != nil {
			log.Fatal(err)
		}
		subscriptions = append(subscriptions, subscription)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return subscriptions
}

func ComposeSubscription(line string) (models.Subscription, error) {
	columnIndex := strings.Index(line, ":")
	dashIndex := strings.Index(line, "-")
	if columnIndex == -1 || dashIndex == -1 {
		emptyEmail := *models.NewEmailAddress("")
		emptyPair := *models.NewCurrencyPair("", "")
		emptySubscription := *models.NewSubscription(emptyEmail, emptyPair)
		return emptySubscription, errors.ErrCouldNotMarshallSubscription
	}
	email := models.NewEmailAddress(line[:columnIndex])
	base := line[columnIndex+1 : dashIndex]
	quote := line[dashIndex+1:]
	pair := models.NewCurrencyPair(base, quote)
	return *models.NewSubscription(*email, *pair), nil
}
