package storage

import (
	"GenesisTask/config"
	"GenesisTask/pkg/application"
	"GenesisTask/pkg/domain/models"
	"bufio"
	"log"
	"os"
)

type EmailFileRepository struct {
	path string
}

func NewFileRepository() application.EmailAddressesStorage {
	path := config.Get().StorageFile
	return &EmailFileRepository{path}
}

func (r *EmailFileRepository) IsSaved(email models.EmailAddress) bool {
	file, err := os.OpenFile(r.path, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if email.GetAddress() == scanner.Text() {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return false
}

func (r *EmailFileRepository) AddEmail(email models.EmailAddress) error {
	file, err := os.OpenFile(r.path, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(email.GetAddress() + "\n")
	if err != nil {
		return err
	}
	err = file.Sync()
	if err != nil {
		return err
	}

	return err
}

func (r *EmailFileRepository) GetAll() []models.EmailAddress {
	file, err := os.OpenFile(r.path, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var addresses []models.EmailAddress
	for scanner.Scan() {
		addresses = append(addresses, *models.NewEmailAddress(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return addresses
}
