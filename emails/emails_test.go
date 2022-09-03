package emails

import (
	"GenesisTask/config"
	"os"
	"strings"
	"testing"
)

func TestEnsureFileExists(t *testing.T) {
	EnsureFileExists()

	path := config.Get().StorageFile
	defer func() {
		directory := strings.Split(path, "/")[0]
		os.RemoveAll(directory)
	}()
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		t.Fail()
	}
}

func TestEnsureDirectoryExists(t *testing.T) {
	ensureDirectoryExists()

	path := config.Get().StorageFile
	directory := strings.Split(path, "/")[0]
	defer func() {
		os.RemoveAll(directory)
	}()
	_, err := os.Stat(directory)
	if os.IsNotExist(err) {
		t.Fail()
	}
}

func TestThatEmailIsAdded(t *testing.T) {
	EnsureFileExists()
	emailToAdd := "plainemail@gmail.com"

	AddEmail(emailToAdd)

	defer func() {
		path := config.Get().StorageFile
		directory := strings.Split(path, "/")[0]
		os.RemoveAll(directory)
	}()
	if !IsEmailSaved(emailToAdd) {
		t.Fail()
	}
}

func TestThatEmailRecordCheckIsValid(t *testing.T) {
	EnsureFileExists()
	emailToAdd := "plainemail@gmail.com"
	anotherEmail := "notinfile@gmail.com"
	AddEmail(emailToAdd)
	defer func() {
		path := config.Get().StorageFile
		directory := strings.Split(path, "/")[0]
		os.RemoveAll(directory)
	}()

	if !IsEmailSaved(emailToAdd) {
		t.Fail()
	}
	if IsEmailSaved(anotherEmail) {
		t.Fail()
	}
}
