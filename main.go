package main

import (
	db "example.com/rest-api/db"
	routes "example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	appServer := gin.Default()
	routes.RegisterRoutes(appServer)

	appServer.Run(":8000") //listens localhost:8000
}
