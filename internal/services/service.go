package services

import (
	"web-todo-service/internal/models"
	"web-todo-service/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list models.TodoList) (int, error)
	GetAll(userId int) (*[]models.TodoList, error)
	GetById(userId int, listId int) (*models.TodoList, error)
	Update(userId int, list models.TodoList) error
	Delete(listId int) error
}

type TodoItem interface {
}
type Service struct {
	Authorization
	TodoList
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
	}
}
