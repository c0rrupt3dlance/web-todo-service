package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"log"
	"strings"
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
		logrus.Printf("error: %s", err)
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

func (r *TodoItemPostgres) GetById(userId, itemId int) (models.TodoItem, error) {
	query := fmt.Sprintf(`select ti.id, ti.title, ti.description, ti.done from %s ti 
                                                inner join %s li on li.item_id = ti.id 
                                                    inner join %s ul on ul.list_id = li.list_id 
                                                        where ti.id=$1 and ul.user_id=$2`,
		todoItemTable, listsItemsTable, usersListsTable)

	row := r.pool.QueryRow(context.Background(), query, itemId, userId)
	var item models.TodoItem
	if err := row.Scan(&item.Id, &item.Title, &item.Description, &item.Done); err != nil {
		logrus.Printf("error: %s", err)
		return item, err
	}

	return item, nil
}

func (r *TodoItemPostgres) Update(userId, itemId int, inputItem models.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argsId := 1

	if inputItem.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argsId))
		args = append(args, *inputItem.Title)
		argsId++
	}

	if inputItem.Description != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argsId))
		args = append(args, *inputItem.Description)
		argsId++
	}

	if inputItem.Done != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argsId))
		args = append(args, *inputItem.Done)
		argsId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`update %s ti set %s from %s li, %s ul 
                    where ti.id = li.item_id and li.list_id = ul.list_id and ul.user_id = $%d and li.item_id = $%d`,
		todoItemTable, setQuery, listsItemsTable, usersListsTable, argsId, argsId+1)

	args = append(args, userId, itemId)

	_, err := r.pool.Exec(context.Background(), query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *TodoItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf(`delete from %s ti using %s li, %s ul 
       where ti.id = li.item_id and li.list_id = ul.list_id and ul.user_id=$1 and ti.id =  $2`,
		todoItemTable, listsItemsTable, usersListsTable)

	_, err := r.pool.Exec(context.Background(), query, userId, itemId)
	if err != nil {
		logrus.Printf("sql error: %s", err)
		return err
	}

	return nil
}
