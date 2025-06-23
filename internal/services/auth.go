package services

import (
	"crypto/sha1"
	"fmt"
	"web-todo-service/internal/models"
	"web-todo-service/internal/repository"
)

const salt string = "3k2mr20g4.!MR#"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (a *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = a.generatePasswordHash(user.Password)
	return a.repo.CreateUser(user)
}

func (a *AuthService) GetUser(username, password string) (models.User, error) {
	return models.User{}, nil
}

func (a *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
