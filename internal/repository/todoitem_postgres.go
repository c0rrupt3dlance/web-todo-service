package repository

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"web-todo-service/internal/models"
)

type TodoItemPostgres struct {
	pool *pgxpool.Pool
}

func NewTodoItemPostgres(pool *pgxpool.Pool) *TodoItemPostgres {
	return &TodoItemPostgres{pool: pool}
}

func (r *TodoItemPostgres) Create(userId int, listId int, item models.TodoItem) (int, error) {
}

func (r *TodoItemPostgres) GetAll(userId, listId int) (*[]models.TodoItem, error) {
	return nil, nil
}
