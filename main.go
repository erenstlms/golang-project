package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Note struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var notes = []Note{
	{ID: "1", Title: "İlk Not", Content: "Go harika bir dil!"},
	{ID: "2", Title: "İkinci Not", Content: "REST API basit."},
}

func getNotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, notes)
}

func getNoteByID(c *gin.Context) {
	id := c.Param("id")
	for _, n := range notes {
		if n.ID == id {
			c.IndentedJSON(http.StatusOK, n)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "not bulunamadı"})
}

func createNote(c *gin.Context) {
	var newNote Note
	if err := c.BindJSON(&newNote); err != nil {
		return
	}
	notes = append(notes, newNote)
	c.IndentedJSON(http.StatusCreated, newNote)
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hoşgeldiniz! API çalışıyor."})
	})
	router.GET("/notes", getNotes)
	router.GET("/notes/:id", getNoteByID)
	router.POST("/notes", createNote)

	router.Run(":8080")
}
