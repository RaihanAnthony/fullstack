package models

import (
	"context"
	"database/sql"

	"github.com/Raihanpoke/fullstack/config"
	"github.com/Raihanpoke/fullstack/entities"
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

// func (um UserModel) Where(user entities.User) (int64, error) {
// 	ctx := context.Background()

// 	query := "SELECT *FROM users WHERE email = ?, password = ?"
// 	result, err := um.db.ExecContext(ctx, query, user.Email, user.Password)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func (um UserModel) Create(user entities.User) (int64, error) {
	ctx := context.Background()

	query := "INSERT INTO users(id, first_name, last_name, email, password, phone, created_at, updated_at) VALUE(?,?,?,?,?,?,?,?)"
	result, err := um.db.ExecContext(ctx, query, user.ID, user.First_Name, user.Last_Name, user.Email, user.Password, user.Phone, user.Created_At, user.Updated_At)
	if err != nil {
		return 0, err
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId, nil
}
