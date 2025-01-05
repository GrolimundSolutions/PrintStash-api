package handlers

import (
    "net/http"
    "strconv"

    "github.com/GrolimundSolutions/PrintStash-api/database"
    "github.com/GrolimundSolutions/PrintStash-api/models"
    "github.com/gin-gonic/gin"
)

// GetAllColors returns all colors
func GetAllColors(c *gin.Context) {
    var colors []models.Color
    db := database.GetDB(c.Request.Context())
    if result := db.Find(&colors); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving colors"})
        return
    }
    c.JSON(http.StatusOK, colors)
}

// GetColor returns a single color
func GetColor(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 16)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    var color models.Color
    db := database.GetDB(c.Request.Context())
    if result := db.First(&color, int16(id)); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Color not found"})
        return
    }

    c.JSON(http.StatusOK, color)
}

// CreateColor creates a new color
func CreateColor(c *gin.Context) {
    var color models.Color
    if err := c.ShouldBindJSON(&color); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db := database.GetDB(c.Request.Context())
    if result := db.Create(&color); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating color"})
        return
    }

    c.JSON(http.StatusCreated, color)
}

// UpdateColor updates an existing color
func UpdateColor(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 16)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    var color models.Color
    db := database.GetDB(c.Request.Context())
    if result := db.First(&color, int16(id)); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Color not found"})
        return
    }

    if err := c.ShouldBindJSON(&color); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if result := db.Save(&color); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating color"})
        return
    }

    c.JSON(http.StatusOK, color)
}

// DeleteColor deletes a color
func DeleteColor(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    db := database.GetDB(c.Request.Context())
    if result := db.Delete(&models.Color{}, uint(id)); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting color"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Color deleted successfully"})
}