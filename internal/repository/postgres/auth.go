package postgres

import (
	"HomeService/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) Register(user model.UserRegister) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (email, user_type, password) VALUES($1,$2,$3) RETURNING id", userTable)
	err := r.db.QueryRow(query, user.Email, user.UserType, user.Password).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *AuthPostgres) Login(user model.UserLogin) (string, error) {
	var user_type string
	query := fmt.Sprintf("SELECT user_type FROM %s WHERE id=$1 AND password=$2", userTable)
	err := r.db.QueryRow(query, user.Id, user.Password).Scan(&user_type)
	if err != nil {
		return "", err
	}
	return user_type, nil
} //UserLogin
