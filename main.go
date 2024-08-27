package main

import (
	"GOTODO/database"
	"GOTODO/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	//fmt.Println("hii")
	Router := gin.Default()
	database.ConnectDatabase()

	//Router.Use(cors.Default())
	Router.POST("/addtodo", handlers.CreateTodo)
	Router.GET("/", handlers.GetTodos)
	Router.Run("localhost:3000")
}
