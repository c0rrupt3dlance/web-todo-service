package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"web-todo-service/internal/models"
)

type AuthPostgres struct {
	pool *pgxpool.Pool
}

func NewAuthPostgres(pool *pgxpool.Pool) Authorization {
	return &AuthPostgres{pool: pool}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s(name, username, password_hash) values ($1, $2, $3) returning id", usersTable)
	row := r.pool.QueryRow(context.Background(), query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		log.Printf("error wuth db query: %s", err)
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("select id from %s where username = $1 and password_hash = $2", usersTable)
	row := r.pool.QueryRow(context.Background(), query, username, password)
	if err := row.Scan(&user.Id); err != nil {
		log.Printf("sql error: %s", err)
		return models.User{}, err
	}

	return user, nil
}
