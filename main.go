package main

import (
	"lilurl/api"
	"lilurl/database"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbInstance := database.CreateConnection()
	api.StartHttpServer(dbInstance)
}
