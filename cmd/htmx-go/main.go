package main

import (
	"log"
	"net/http"

	"github.com/Jayanth-Parthsarathy/htmx-go/config"
	"github.com/Jayanth-Parthsarathy/htmx-go/database"
	"github.com/Jayanth-Parthsarathy/htmx-go/internal/renderer"
	"github.com/Jayanth-Parthsarathy/htmx-go/views"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := database.Init()
	r.HTMLRender = renderer.Default
	r.Static("/static", "./static")
	r.SetTrustedProxies(nil)
	r.GET("/", func(c *gin.Context) {
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
	})
	r.POST("/new-note", func(c *gin.Context) {
		title := c.PostForm("title")
		content := c.PostForm("content")
		_, err := db.Exec("INSERT INTO notes (title, content) VALUES ($1, $2)", title, content)
		if err != nil {
			log.Fatal(err)
		}
		c.HTML(http.StatusOK, "", views.Note(title, content))
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
