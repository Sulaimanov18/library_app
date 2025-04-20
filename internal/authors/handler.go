package authors

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterAuthorRoutes(r *gin.Engine, service *AuthorService) {
	r.GET("/authors", func(c *gin.Context) {
		c.JSON(http.StatusOK, service.GetAll())
	})

	r.GET("/authors/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		author, err := service.GetByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
			return
		}
		c.JSON(http.StatusOK, author)
	})

	r.POST("/authors", func(c *gin.Context) {
		var req CreateAuthorRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		author := service.Create(req)
		c.JSON(http.StatusCreated, author)
	})
}
