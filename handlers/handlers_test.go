// handlers_test.go
package handlers_test

import (
	"bytes"
	"encoding/json"
	"GOTODO/database"
	"GOTODO/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHandlers(t *testing.T) {
	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Handlers Suite")
}

var _ = ginkgo.Describe("Todo Handlers", func() {
	var router *gin.Engine

	ginkgo.BeforeEach(func() {
		// Initialize database connection
		database.ConnectDatabase()

		// Set up Gin router
		router = gin.Default()

		// Register handlers
		router.GET("/todos", handlers.GetTodos)
		router.POST("/todos", handlers.CreateTodo)
		router.GET("/todos/:id", handlers.Findtodo)
		router.DELETE("/todos/:id", handlers.Deletetodo)
	})

	ginkgo.Context("GET /todos", func() {
		ginkgo.It("should return an empty list of todos", func() {
			req, _ := http.NewRequest("GET", "/todos", nil)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusOK))
			Expect(resp.Body.String()).To(Equal("[]"))
		})
	})

	ginkgo.Context("POST /todos", func() {
		ginkgo.It("should create a new todo", func() {
			todo := map[string]interface{}{
				"title": "New Todo",
			}
			jsonData, _ := json.Marshal(todo)

			req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusCreated))
			Expect(resp.Body.String()).To(ContainSubstring("New Todo"))
		})
	})
})
