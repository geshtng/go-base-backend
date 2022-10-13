package main

import (
	"github.com/gin-gonic/gin"

	"github.com/geshtng/go-base-backend/config"
	"github.com/geshtng/go-base-backend/db"
	"github.com/geshtng/go-base-backend/internal/handlers"
	"github.com/geshtng/go-base-backend/internal/routes"
)

func main() {
	port := config.InitConfigPort()

	r := gin.Default()

	err := db.Connect()
	if err != nil {
		panic(err)
	}

	h := handlers.InitAllHandlers()

	routes.InitAllRoutes(r, h)

	err = r.Run(port)
	if err != nil {
		panic(err)
	}
}
