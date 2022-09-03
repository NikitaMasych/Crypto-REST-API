package emails

import (
	"GenesisTask/config"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

func TestEnsureFileExists(t *testing.T) {
	err := godotenv.Load("./../.env") // from upper directory
	if err != nil {
		t.Error(err)
	}

	EnsureFileExists()

	path := config.Get().StorageFile
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		t.Fail()
	}
}

func TestEnsureDirectoryExists(t *testing.T) {
	err := godotenv.Load("./../.env") // from upper directory
	if err != nil {
		t.Error(err)
	}

	ensureDirectoryExists()

	path := config.Get().StorageFile
	directory := strings.Split(path, "/")[0]
	_, err = os.Stat(directory)
	if os.IsNotExist(err) {
		t.Fail()
	}
}

func TestAddEmail(t *testing.T) {
	err := godotenv.Load("./../.env") // from upper directory
	if err != nil {
		t.Error(err)
	}
	EnsureFileExists()
	emailToAdd := "plainemail@gmail.com"

	AddEmail(emailToAdd)

	if !IsEmailSaved(emailToAdd) {
		t.Fail()
	}
}

func TestIsEmailSaved(t *testing.T) {
	err := godotenv.Load("./../.env") // from upper directory
	if err != nil {
		t.Error(err)
	}
	EnsureFileExists()
	emailToAdd := "plainemail@gmail.com"
	anotherEmail := "notinfile@gmail.com"
	AddEmail(emailToAdd)

	if !IsEmailSaved(emailToAdd) {
		t.Fail()
	}
	if IsEmailSaved(anotherEmail) {
		t.Fail()
	}
}
