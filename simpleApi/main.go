package main

import (
	"database/sql"
	"log"
	"os"
	"test/albums"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// Mock data, changed for seeding db
// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

func main() {
	// Using gin, more preferable net/http or chi
	g := gin.Default()

	file, err := os.Create("sqlite-db.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("Database created")

	sqlitedb, err := sql.Open("sqlite3", "./sqlite-db.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	sqlitedb.SetMaxOpenConns(1)
	defer sqlitedb.Close()
	createTable(sqlitedb)

	g.GET("/albums", albums.GetAlbums(sqlitedb))
	g.POST("/albums", albums.PostAlbums(sqlitedb))
	g.GET("/albums/:id", albums.GetAlbum(sqlitedb))

	g.Run()
}

func createTable(db *sql.DB) {
	createAlbumTableSQL := `CREATE TABLE IF NOT EXISTS albums (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		artist TEXT,
		price FLOAT
	);`

	log.Println("Create album table")
	statement, err := db.Prepare(createAlbumTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Table albums created")

	log.Println("Seed album data")
	insertSeedDataSQL1 := `INSERT OR IGNORE INTO albums (title, artist, price) VALUES ('Blue Train', 'John Coltrane', 56.99);`
	insertSeedDataSQL2 := `INSERT OR IGNORE INTO albums (title, artist, price) VALUES ('Jeru', 'Gerry Mulligan', 17.99);`

	insert, err := db.Prepare(insertSeedDataSQL1)
	if err != nil {
		log.Fatal(err.Error())
	}
	insert.Exec()

	insert2, err := db.Prepare(insertSeedDataSQL2)
	if err != nil {
		log.Fatal(err.Error())
	}
	insert2.Exec()

	log.Println("Data seeded")

}
