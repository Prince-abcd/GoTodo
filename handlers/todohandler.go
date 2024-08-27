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
func Findtodo(c *gin.Context) {
	id := c.Param("id")

	var todo database.Todo

	if result := database.DB.First(&todo, id); result.Error != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}
func Deletetodo(c *gin.Context) {
	var todo database.Todo
	id := c.Param("id")
	if err := database.DB.Where("Id = ?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	database.DB.Delete(&todo)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
