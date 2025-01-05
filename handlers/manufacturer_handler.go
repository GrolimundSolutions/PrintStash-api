package handlers

import (
	"net/http"
	"strconv"

	"github.com/GrolimundSolutions/PrintStash-api/database"
	"github.com/GrolimundSolutions/PrintStash-api/models"
	"github.com/gin-gonic/gin"
)

// GetAllManufacturers returns all manufacturers
func GetAllManufacturers(c *gin.Context) {
	var manufacturers []models.Manufacturer
	db := database.GetDB(c.Request.Context())
	if result := db.Find(&manufacturers); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving manufacturers"})
		return
	}
	c.JSON(http.StatusOK, manufacturers)
}

// GetManufacturer returns a single manufacturer
func GetManufacturer(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var manufacturer models.Manufacturer
	db := database.GetDB(c.Request.Context())
	if result := db.First(&manufacturer, int16(id)); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Manufacturer not found"})
		return
	}

	c.JSON(http.StatusOK, manufacturer)
}

// CreateManufacturer creates a new manufacturer
func CreateManufacturer(c *gin.Context) {
	var manufacturer models.Manufacturer
	if err := c.ShouldBindJSON(&manufacturer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB(c.Request.Context())
	if result := db.Create(&manufacturer); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating manufacturer"})
		return
	}

	c.JSON(http.StatusCreated, manufacturer)
}

// UpdateManufacturer updates an existing manufacturer
func UpdateManufacturer(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var manufacturer models.Manufacturer
	db := database.GetDB(c.Request.Context())
	if result := db.First(&manufacturer, uint(id)); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Manufacturer not found"})
		return
	}

	if err := c.ShouldBindJSON(&manufacturer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Save(&manufacturer); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating manufacturer"})
		return
	}

	c.JSON(http.StatusOK, manufacturer)
}

// DeleteManufacturer deletes a manufacturer
func DeleteManufacturer(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	db := database.GetDB(c.Request.Context())
	if result := db.Delete(&models.Manufacturer{}, uint(id)); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting manufacturer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Manufacturer deleted successfully"})
}
