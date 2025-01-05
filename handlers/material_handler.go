package handlers

import (
    "net/http"
    "strconv"

    "github.com/GrolimundSolutions/PrintStash-api/database"
    "github.com/GrolimundSolutions/PrintStash-api/models"
    "github.com/gin-gonic/gin"
)

// GetAllMaterials returns all materials
func GetAllMaterials(c *gin.Context) {
    var materials []models.Material
    db := database.GetDB(c.Request.Context())
    if result := db.Find(&materials); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving materials"})
        return
    }
    c.JSON(http.StatusOK, materials)
}

// GetMaterial returns a single material
func GetMaterial(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 16)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    var material models.Material
    db := database.GetDB(c.Request.Context())
    if result := db.First(&material, int16(id)); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Material not found"})
        return
    }

    c.JSON(http.StatusOK, material)
}

// CreateMaterial creates a new material
func CreateMaterial(c *gin.Context) {
    var material models.Material
    if err := c.ShouldBindJSON(&material); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db := database.GetDB(c.Request.Context())
    if result := db.Create(&material); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating material"})
        return
    }

    c.JSON(http.StatusCreated, material)
}

// UpdateMaterial updates an existing material
func UpdateMaterial(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    var material models.Material
    db := database.GetDB(c.Request.Context())
    if result := db.First(&material, uint(id)); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Material not found"})
        return
    }

    if err := c.ShouldBindJSON(&material); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if result := db.Save(&material); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating material"})
        return
    }

    c.JSON(http.StatusOK, material)
}

// DeleteMaterial deletes a material
func DeleteMaterial(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 16)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    db := database.GetDB(c.Request.Context())
    if result := db.Delete(&models.Material{}, int16(id)); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting material"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Material deleted successfully"})
}