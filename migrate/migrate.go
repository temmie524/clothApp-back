package main

import (
	"clothApp-backend/db"
	"clothApp-backend/model"
	"fmt"
	"os"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully migrated!!")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(
		model.User{},
		model.Brand{},
		model.Category{},
		model.Comment{},
		model.Item{},
		model.Like{},
		model.Size{},
		model.Season{},
	)
	fmt.Println(os.Getenv("DB_NAME"))
}
