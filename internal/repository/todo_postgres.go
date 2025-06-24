package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"web-todo-service/internal/models"
)

type TodoListPostgres struct {
	pool *pgxpool.Pool
}

func NewTodoListPostgres(pool *pgxpool.Pool) *TodoListPostgres {
	return &TodoListPostgres{
		pool: pool,
	}
}

func (r *TodoListPostgres) Create(userId int, list models.TodoList) (int, error) {
	var listId int
	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return 0, err
	}
	createListQuery := fmt.Sprintf("insert into %s (title, description) values ($1, $2) returning id", todoListsTable)
	row := tx.QueryRow(context.Background(), createListQuery, list.Title, list.Description)
	if err := row.Scan(&listId); err != nil {
		log.Printf("got error during inserting new list: %s", err)
		tx.Rollback(context.Background())
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("insert into %s (user_id, list_id) values ($1, $2)", usersListsTable)

	_, err = tx.Exec(context.Background(), createUsersListQuery, userId, listId)
	if err != nil {
		log.Printf("got error during inserting new users_list: %s", err)
		tx.Rollback(context.Background())
		return 0, err
	}

	tx.Commit(context.Background())
	return listId, nil
}
