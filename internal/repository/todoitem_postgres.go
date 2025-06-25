package repository

import (
	"context"
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

func (r *TodoItemPostgres) GetAll(listId int) (*[]models.TodoItem, error) {
	items := make([]models.TodoItem, 0)

	query := fmt.Sprintf("select ti.id, ti.title, ti.description, ti.done from %s ti inner join %s li on ti.id = li.item_id where li.list_id = $1", todoItemTable, listsItemsTable)
	rows, err := r.pool.Query(context.Background(), query, listId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item models.TodoItem
		if err := rows.Scan(&item.Id, &item.Title, &item.Description, &item.Done); err != nil {
			return nil, err
		}
	}
	return &items, nil
}

func (r *TodoItemPostgres) Create(listId int, item *models.TodoItem) (int, error) {
	return 0, nil
}
