package handlers

import (
	"log"
	"net/http"

	"github.com/Jayanth-Parthsarathy/htmx-go/config"
	"github.com/Jayanth-Parthsarathy/htmx-go/database"
	"github.com/Jayanth-Parthsarathy/htmx-go/views"
	"github.com/gin-gonic/gin"
)

func DisplayNotes(c *gin.Context) {
	db := database.GetDB()
	rows, err := db.Query("SELECT title, content FROM notes")
	if err != nil {
		log.Fatal(err)
	}
	var notes []config.Note
	defer rows.Close()
	for rows.Next() {
		var title string
		var content string
		err := rows.Scan(&title, &content)
		if err != nil {
			log.Fatal(err)
		}
		note := config.Note{Title: title, Content: content}
		notes = append(notes, note)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return
	}
	c.HTML(http.StatusOK, "", views.Home(notes))
}

func AddNote(c *gin.Context) {
	db := database.GetDB()
	title := c.PostForm("title")
	content := c.PostForm("content")
	_, err := db.Exec("INSERT INTO notes (title, content) VALUES ($1, $2)", title, content)
	if err != nil {
		log.Fatal(err)
	}
	c.HTML(http.StatusOK, "", views.Note(title, content))
}
