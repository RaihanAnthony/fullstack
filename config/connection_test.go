package config

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Raihanpoke/fullstack/entities"
	"github.com/google/uuid"
)

func TestConn(t *testing.T) {
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var user entities.User
	ctx := context.Background()

	user.ID = uuid.New().String()
	user.First_Name = "raihan"
	user.Last_Name = "malay"
	user.Email = "raihan@gmail.com"
	user.Password = "093982"
	user.Phone = "08953297969"
	user.Created_At = time.Now()
	user.Updated_At = time.Now()

	query := "INSERT INTO users(id, first_name, last_name, email, password, phone, created_at, updated_at) VALUE(?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = db.ExecContext(ctx, query, user.ID, user.First_Name, user.Last_Name, user.Email, user.Password, user.Phone, user.Created_At, user.Updated_At)
	if err != nil {
		panic(err)
	}

	fmt.Println("succesfully insert data to mysql")

}
