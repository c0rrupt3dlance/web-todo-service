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
}

type ListItem interface {
}
type Repository struct {
	Authorization
	TodoList
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(pool),
		TodoList:      NewTodoListPostgres(pool),
	}
}
