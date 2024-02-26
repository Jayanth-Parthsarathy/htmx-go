package main

import (
	"github.com/Jayanth-Parthsarathy/htmx-go/database"
	"github.com/Jayanth-Parthsarathy/htmx-go/internal/router"
)

func main() {
	database.Init()
  router.Init()
}
