package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type DataBase struct {
	Instance *sql.DB
}

func CreateConnection() *DataBase {
	file, err := os.Create("sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	fmt.Println("sqlite-database.db created")

	sqliteDatabase, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.Fatal(err)
	}

	db := &DataBase{Instance: sqliteDatabase}
	db.Migrate()

	return db
}

func (db *DataBase) Migrate() {
	sqlFiles, err := os.ReadDir("./database/migrations")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range sqlFiles {
		slqFile, err := os.ReadFile("./database/migrations/" + file.Name())
		if err != nil {
			log.Fatal(err.Error())
		}

		statement, err := db.Instance.Prepare(string(slqFile))
		if err != nil {
			log.Fatal(err.Error())
		}
		statement.Exec()
	}
}
