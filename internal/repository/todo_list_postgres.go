package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"strings"
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

func (r *TodoListPostgres) GetAll(userId int) (*[]models.TodoList, error) {
	var userLists []models.TodoList
	query := fmt.Sprintf(`select tl.id, tl.title, tl.description from %s tl
        inner join %s ul on tl.id = ul.list_id where ul.user_id = $1`,
		todoListsTable, usersListsTable)

	rows, err := r.pool.Query(context.Background(), query, userId)
	defer rows.Close()
	if err != nil {
		log.Printf("sql error: %s", err)
		return nil, err
	}

	for rows.Next() {
		var list models.TodoList
		if err = rows.Scan(&list.Id, &list.Title, &list.Description); err != nil {
			log.Printf("sql error: %s", err)
			return nil, err
		}

		userLists = append(userLists, list)

	}

	return &userLists, nil
}

func (r *TodoListPostgres) GetById(userId int, listId int) (models.TodoList, error) {
	var list models.TodoList

	query := fmt.Sprintf(`select tl.id, tl.title, tl.description from %s tl 
		inner join %s ul on tl.id = ul.list_id where ul.user_id=$1 and ul.list_id=$2`,
		todoListsTable, usersListsTable)

	row := r.pool.QueryRow(context.Background(), query, userId, listId)

	err := row.Scan(&list.Id, &list.Title, &list.Description)

	return list, err
}

func (r *TodoListPostgres) Update(userId, listId int, inputList models.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if inputList.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *inputList.Title)
		argId++
	}

	if inputList.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *inputList.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("update %s tl set %s from %s ul where tl.id=ul.list_id and ul.list_id = $%d and ul.user_id=$%d",
		todoListsTable, setQuery, usersListsTable, argId, argId+1)

	args = append(args, listId, userId)

	_, err := r.pool.Exec(context.Background(), query, args...)
	log.Println(query)
	if err != nil {
		log.Printf("sql error: %s", err)
		return err
	}

	return nil
}

func (r *TodoListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf(`delete from %s tl using %s ul where tl.id = ul.list_id 
                                  and ul.user_id=$1 and ul.list_id=$2`,
		todoListsTable, usersListsTable)
	_, err := r.pool.Exec(context.Background(), query, userId, listId)
	if err != nil {
		log.Printf("sql error: %s", err)
	}
	return err
}
