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

// Struct tags such as json:"artist" specify what a field’s name should be when the struct’s contents are serialized into JSON.
// Without them, the JSON would use the struct’s capitalized field names – a style not as common in JSON.

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	// serialize the struct into JSON and add it to the response.
	c.IndentedJSON(http.StatusOK, albums)
	// c.JSON(http.StatusOK, albums)
	// Note that you can replace Context.IndentedJSON with a call to Context.JSON to send more compact JSON.
}

// gin.Context is the most important part of Gin. It carries request details,
// validates and serializes JSON, and more.

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	// Use Context.Param to retrieve the id path parameter from the URL.
	// When you map this handler to a path, you’ll include a placeholder for the parameter in the path.

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	// With Gin, you can associate a handler with an HTTP method-and-path combination.
	// In this way, you can separately route requests sent to a single path based on the method the client is using.
	// In Gin, the colon preceding an item in the path signifies that the item is a path parameter.

	router.Run("localhost:8080")
	// Use the Run function to attach the router to an http.Server and start the server.
}
