package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"log"
	"web-todo-service/internal/models"
)

type TodoItemPostgres struct {
	pool *pgxpool.Pool
}

func NewTodoItemPostgres(pool *pgxpool.Pool) *TodoItemPostgres {
	return &TodoItemPostgres{pool: pool}
}

func (r *TodoItemPostgres) Create(listId int, item models.TodoItem) (int, error) {
	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return 0, err
	}
	var itemId int
	createItemQuery := fmt.Sprintf(`insert into %s (title, description) values ($1, $2) returning id`, todoItemTable)

	row := tx.QueryRow(context.Background(), createItemQuery, item.Title, item.Description)
	if err = row.Scan(&itemId); err != nil {
		tx.Rollback(context.Background())
		return 0, err
	}

	listsItemQuery := fmt.Sprintf(`insert into %s (list_id, item_id) values ($1, $2)`, listsItemsTable)
	_, err = tx.Exec(context.Background(), listsItemQuery, listId, itemId)
	if err != nil {
		tx.Rollback(context.Background())
		return 0, err
	}

	return itemId, tx.Commit(context.Background())
}

func (r *TodoItemPostgres) GetAll(userId, listId int) ([]models.TodoItem, error) {
	items := make([]models.TodoItem, 0)
	query := fmt.Sprintf(`select ti.id, ti.title, ti.description, ti.done from %s ti 
                                       inner join %s li on li.item_id = ti.id
                                       inner join %s ul on ul.list_id = li.list_id where li.list_id = $1 and ul.user_id = $2`,
		todoItemTable, listsItemsTable, usersListsTable)

	rows, err := r.pool.Query(context.Background(), query, listId, userId)
	if err != nil {
		logrus.Printf("Error: %s", err)
		return nil, err
	}

	for rows.Next() {
		var item models.TodoItem
		if err = rows.Scan(&item.Id, &item.Title, &item.Description, &item.Done); err != nil {
			log.Printf("sql error: %s", err)
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}
