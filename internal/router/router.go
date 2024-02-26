package router

import (
	"github.com/Jayanth-Parthsarathy/htmx-go/handlers"
	"github.com/Jayanth-Parthsarathy/htmx-go/internal/renderer"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	r.HTMLRender = renderer.Default
	r.SetTrustedProxies(nil)
	r.GET("/", handlers.DisplayNotes)
	r.POST("/new-note", handlers.AddNote)
	r.Run()
}
