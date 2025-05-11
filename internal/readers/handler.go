package readers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterReaderRoutes(r *gin.Engine, service *ReaderService) {
	r.GET("/readers", func(c *gin.Context) {
		readers, err := service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, readers)
	})

	r.GET("/readers/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		reader, err := service.GetByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if reader == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Reader not found"})
			return
		}
		c.JSON(http.StatusOK, reader)
	})

	r.POST("/readers", func(c *gin.Context) {
		var req CreateReaderRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		reader, err := service.Create(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, reader)
	})

	r.PUT("/readers/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var req UpdateReaderRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := service.Update(id, req); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusOK)
	})

	r.DELETE("/readers/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := service.Delete(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	})
}
