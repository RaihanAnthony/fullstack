package config

import (
	"context"
	"fmt"
	"testing"
)

func TestConn(t *testing.T) {
	db, err := DBConenection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctx := context.Background()

	id := "bhbjdsvgvcdhbhj"
	first_name := "raihan"
	last_name := "malay"

	query := "INSERT INTO users(id, first_name, last_name) VALUE(?, ?, ?)"
	_, err = db.ExecContext(ctx, query, id, first_name, last_name)
	if err != nil {
		panic(err)
	}

	fmt.Println("succesfully insert data to mysql")

}
