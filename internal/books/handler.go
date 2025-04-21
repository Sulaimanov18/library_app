package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(r *gin.Engine, service *BookService) {
	r.GET("/books", getAllBooksHandler(service))
	r.POST("/books", createBookHandler(service))
}

// Обработчик GET /books
func getAllBooksHandler(service *BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, books)
	}
}

// Обработчик POST /books
func createBookHandler(service *BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateBookRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		book, err := service.Create(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, book)
	}
}
