// A standalone program (as opposed to a library) is always in package main.
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/album/:id", getAlbumById)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	// Could just use c.JSON() for uglier json
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(c *gin.Context) {
	// Extract query param "id"
	id := c.Param("id")

	// Find the album with the matching id
	for _, album := range albums {
		if album.ID == id {
			// 200 & album returned via context if found inside loop
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	// If album not found, return message in header and 404
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

// post an album
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Binds request "data" to the newAlbum var.
	// If the data is not of the same shape what happens? => Field that is not included is left empty

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
