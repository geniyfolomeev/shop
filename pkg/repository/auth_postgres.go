package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"shop"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user shop.User) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (login, first_name, surname, password_hash) "+
			"values ($1, $2, $3, $4) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Login, user.Name, user.Surname, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	query = fmt.Sprintf(
		"INSERT INTO %s (user_id, dark_theme, role) "+
			"values ($1, $2, $3)", settingsTable)
	_, err := r.db.Exec(query, id, false, "user")
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(login, password string) (shop.User, error) {
	var user shop.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE login=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, login, password)

	return user, err
}
