package emails

import (
	"GenesisTask/config"
	"GenesisTask/errors"
	"bufio"
	"log"
	"os"
)

func EnsureFileExists() {
	path := config.Get().StorageFile
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}
}

func AddEmail(email string) error {
	path := config.Get().StorageFile
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if isEmailSaved(file, email) {
		return errors.ErrAlreadyExists
	}

	_, err = file.WriteString(email + "\n")
	if err != nil {
		log.Fatal(err)
	}

	err = file.Sync()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func isEmailSaved(file *os.File, email string) bool {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if email == scanner.Text() {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return false
}
