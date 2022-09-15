package platform

import (
	"log"
	"os"
	"strings"
)

func EnsureFileExists(path string) {
	ensureDirectoryExists(path)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}
}

func ensureDirectoryExists(path string) {
	directory := strings.Split(path, "/")[0]
	_, err := os.Stat(directory)
	if os.IsNotExist(err) {
		err = os.Mkdir(directory, 0700)
		if err != nil {
			log.Fatal(err)
		}
	}
}
