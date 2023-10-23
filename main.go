package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wasimkhandot01/assessment/routes"
)

func main() {
	router := gin.Default()

	//routes
	routes.TasksRoute(router)

	if err := router.Run("localhost:6000"); err != nil {
		panic(err)
	}
}
