package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"web-todo-service/internal/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(Username, Password string) (models.User, error)
}

type TodoList interface {
	Create(userId int, list models.TodoList) (int, error)
	GetAll(userId int) (*[]models.TodoList, error)
	GetById(userId int, listId int) (models.TodoList, error)
	Update(userId int, listId int, list models.UpdateListInput) error
	Delete(userId int, listId int) error
}

type TodoItem interface {
	Create(listId int, item models.TodoItem) (int, error)
	GetAll(userId int, listId int) ([]models.TodoItem, error)
	GetById(userId, itemId int) (models.TodoItem, error)
	Update(userId, itemId int, inputItem models.UpdateItemInput) error
	Delete(userId, itemId int) error
}
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(pool),
		TodoList:      NewTodoListPostgres(pool),
		TodoItem:      NewTodoItemPostgres(pool),
	}
}
