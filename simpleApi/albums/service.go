package albums

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAlbums(db *sql.DB) gin.HandlerFunc {
	return func(g *gin.Context) {
		rows, err := db.Query("SELECT id, title, artist, price FROM albums")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var newAlbums []album
		for rows.Next() {
			var a album
			if err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
				log.Fatal(err.Error())
			}
			newAlbums = append(newAlbums, a)
		}
		g.IndentedJSON(http.StatusOK, newAlbums)
	}
}

func PostAlbums(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var newAlbum album

		if err := c.BindJSON(&newAlbum); err != nil {
			return
		}

		insertSeedDataSQL := `INSERT OR IGNORE INTO albums (title, artist, price) VALUES (` + newAlbum.Title + `, ` + newAlbum.Artist + `, ` + strconv.FormatFloat(newAlbum.Price, 'f', -1, 64) + `);`

		insert, err := db.Prepare(insertSeedDataSQL)
		if err != nil {
			log.Fatal(err.Error())
		}
		insert.Exec()

		c.IndentedJSON(http.StatusCreated, newAlbum)

	}
}

func GetAlbum(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var newAlbum album

		err := db.QueryRow("SELECT id, title, artist, price FROM albums WHERE id = ?", id).Scan(&newAlbum.ID, &newAlbum.Artist, &newAlbum.Title, &newAlbum.Price)
		if err != nil {
			log.Fatal(err)
		}

		c.IndentedJSON(http.StatusOK, newAlbum)
	}
	//id := g.Param("id")

	// for i, a := range albums {
	// 	fmt.Println(i, " ", a)
	// 	if a.ID == id {
	// 		g.IndentedJSON(http.StatusOK, a)
	// 		return
	// 	}
	// }
	// g.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}
