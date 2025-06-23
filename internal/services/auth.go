package services

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
	"web-todo-service/internal/models"
	"web-todo-service/internal/repository"
)

const (
	salt       string = "3k2mr20g4.!MR#"
	signingKey        = "!<RM323rt23!Mi231mer132?t"
	TokenTTL          = 6 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

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

func (a *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := a.repo.GetUser(username, a.generatePasswordHash(password))
	if err != nil {
		return "error", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	tokenStr, err := token.SignedString([]byte(signingKey))
	if err != nil {
		log.Printf("error during token generation: %s", err)
		return "", err
	}

	return tokenStr, nil
}

func (a *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
