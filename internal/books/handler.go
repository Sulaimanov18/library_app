package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(r *gin.Engine, service *BookService) {
	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, service.GetAll())
	})

	r.POST("/books", func(c *gin.Context) {
		var req CreateBookRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		book, err := service.Create(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, book)
	})
}