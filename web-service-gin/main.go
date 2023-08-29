package main

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Vete", Artist: "Bad Bunny", Price: 72.35},
	{ID: "2", Title: "Soy Peor", Artist: "Bad Bunny", Price: 53.12},
	{ID: "3", Title: "Infeliz", Artist: "Arcángel ft. Bad Bunny", Price: 64.89},
	{ID: "4", Title: "Callaíta", Artist: "Bad Bunny ft. Tainy", Price: 39.75},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums/add", postAlbums)
	router.DELETE("/albums/delete/:id", deleteAlbumById)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "album not found"})
}

func deleteAlbumById(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			indexToRemove := id
			id2, err := strconv.ParseInt(indexToRemove, 10, 64)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			albums = append(albums[:id2-1], albums[id2:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"Cancion eliminada": "Cancion " + a.Title + " eliminada correctamente"})
		}

	}
}
