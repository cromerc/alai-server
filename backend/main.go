package main

import (
	"backend/database"
	"backend/routes"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		print(err.Error())
	}

	allArgs := os.Args[1:]

	dropDB := false
	migrateDB := false
	serve := false
	dbOperation := false

	for _, element := range allArgs {
		if element == "drop" {
			dropDB = true
			dbOperation = true
		}

		if element == "migrate" {
			migrateDB = true
			dbOperation = true
		}

		if element == "serve" {
			serve = true
		}
	}

	var gdb *gorm.DB
	if dbOperation {
		fmt.Print("Connecting to database ... ")
		gdb = database.Connect()
		fmt.Println("DONE")
	}

	if dropDB {
		fmt.Print("Dropping database ... ")
		database.DropAll(gdb)
		fmt.Println("DONE")
	}

	if migrateDB {
		fmt.Print("AutoMigrating database ... ")
		database.AutoMigrate(gdb)
		fmt.Println("DONE")
	}

	if dbOperation {
		fmt.Print("Closing database ... ")
		database.Close(gdb)
		fmt.Println("DONE")
	}

	if serve {
		router := routes.Initialize()
		fmt.Println("Serving routes ... DONE")
		routes.Serve(router)
	}
}
