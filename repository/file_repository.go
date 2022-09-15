package repository

import (
	"GenesisTask/config"
	"GenesisTask/model"
	"bufio"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func AttachRepository(r UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("userRepo", r)
		c.Next()
	}
}

type UserFileRepository struct {
	path string
}

func New() *UserFileRepository {
	path := config.Get().StorageFile
	return &UserFileRepository{path: path}
}

func (r *UserFileRepository) IsExist(user *model.User) bool {
	file, err := os.OpenFile(r.path, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if user.GetEmail() == scanner.Text() {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return false
}

func (r *UserFileRepository) Add(user *model.User) error {
	file, err := os.OpenFile(r.path, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(user.GetEmail() + "\n")
	if err != nil {
		return err
	}
	err = file.Sync()
	if err != nil {
		return err
	}

	return err
}

func (r *UserFileRepository) GetUsers() *[]model.User {
	file, err := os.OpenFile(r.path, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var users []model.User
	for scanner.Scan() {
		users = append(users, *model.NewUser(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &users
}
