package main

import (
	"database/sql"
	"fmt"
	"lilurl/api"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var ShortnedUrls = make(map[string]string)

func main() {
	file, err := os.Create("sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	fmt.Println("sqlite-database.db created")
	sqliteDatabase, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	createTableSQL := `CREATE TABLE urls (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"shortId" TEXT,
		"url" TEXT	
	  );`
	statement, err := sqliteDatabase.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	api.StartHttpServer(sqliteDatabase)
}
