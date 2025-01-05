package handlers

import (
	"net/http"
	"strconv"

	"github.com/GrolimundSolutions/PrintStash-api/database"
	"github.com/GrolimundSolutions/PrintStash-api/models"
	"github.com/gin-gonic/gin"
)

// GetAllPrintSettings returns all print settings
func GetAllPrintSettings(c *gin.Context) {
	var printSettings []models.PrintSetting
	db := database.GetDB(c.Request.Context())
	if result := db.Preload("FilamentSpool").Find(&printSettings); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving print settings"})
		return
	}
	c.JSON(http.StatusOK, printSettings)
}

// GetPrintSetting returns a single print setting
func GetPrintSetting(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
	    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
	    return
	}

	var printSetting models.PrintSetting
	db := database.GetDB(c.Request.Context())
	if result := db.Preload("FilamentSpool").First(&printSetting, int32(id)); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Print setting not found"})
		return
	}

	c.JSON(http.StatusOK, printSetting)
}

// CreatePrintSetting creates a new print setting
func CreatePrintSetting(c *gin.Context) {
	var printSetting models.PrintSetting
	if err := c.ShouldBindJSON(&printSetting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB(c.Request.Context())
	if result := db.Create(&printSetting); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating print setting"})
		return
	}

	c.JSON(http.StatusCreated, printSetting)
}

// UpdatePrintSetting updates an existing print setting
func UpdatePrintSetting(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
	    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
	    return
	}

	var printSetting models.PrintSetting
	db := database.GetDB(c.Request.Context())
	if result := db.First(&printSetting, int32(id)); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Print setting not found"})
		return
	}

	if err := c.ShouldBindJSON(&printSetting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Save(&printSetting); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating print setting"})
		return
	}

	c.JSON(http.StatusOK, printSetting)
}

// DeletePrintSetting deletes a print setting
func DeletePrintSetting(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	db := database.GetDB(c.Request.Context())
	if result := db.Delete(&models.PrintSetting{}, uint(id)); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting print setting"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Print setting deleted successfully"})
}
