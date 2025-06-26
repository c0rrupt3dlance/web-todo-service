package services

import (
	"web-todo-service/internal/models"
	"web-todo-service/internal/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{
		repo: repo,
	}
}

func (s *TodoListService) Create(userId int, list models.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) (*[]models.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId int, listId int) (*models.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) Update(userId int, listId int, list models.TodoList) error {
	return s.repo.Update(userId, listId, list)
}

func (s *TodoListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}
