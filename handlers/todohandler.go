package handlers

import (
	"GOTODO/database"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []database.Todo
	result := database.DB.Find(&todos)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}


func CreateTodo(c *gin.Context) {
	var todo database.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := database.DB.Create(&todo)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, todo)
}
