package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/Raihanpoke/fullstack/config"
)

type UserModel struct {
	db *sql.DB
}

func NewUserModel() *UserModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &UserModel{
		db: conn,
	}
}

func (u UserModel) Create(user entities.user) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	query := "INSERT INTO users(first_name, last_name, email, password, phone) VALUE(?,?,?,?,?)"
	result, err := u.db.ExecContext(ctx, query, user.First_Name, user.Last_Name, user.Email, user.Password, user.Phone)
	if err != nil {
		return 0, err
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId, nil
}
