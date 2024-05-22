package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) *sql.DB {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		if _, err := os.Create(filepath); err != nil {
			return nil
		}
	}

	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS heroes (
        "id" INTEGER PRIMARY KEY AUTOINCREMENT,
        "name" TEXT,
        "damage" INTEGER,
        "health" INTEGER,
        "gender" BOOLEAN,
        "class" INTEGER
    );`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	log.Println("Heroes DB connected")
	return db
}
