package handlers

import (
    "net/http"
    "strconv"

    "github.com/GrolimundSolutions/PrintStash-api/database"
    "github.com/GrolimundSolutions/PrintStash-api/models"
    "github.com/gin-gonic/gin"
)

// GetAllFilamentSpools returns all filament spools
func GetAllFilamentSpools(c *gin.Context) {
    var filamentSpools []models.FilamentSpool
    db := database.GetDB(c.Request.Context())
    if result := db.Preload("Manufacturer").Preload("Material").Preload("Color").Find(&filamentSpools); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving filament spools"})
        return
    }
    c.JSON(http.StatusOK, filamentSpools)
}

// GetFilamentSpool returns a single filament spool
func GetFilamentSpool(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    var filamentSpool models.FilamentSpool
    db := database.GetDB(c.Request.Context())
    if result := db.Preload("Manufacturer").Preload("Material").Preload("Color").First(&filamentSpool, int32(id)); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Filament spool not found"})
        return
    }

    c.JSON(http.StatusOK, filamentSpool)
}

// CreateFilamentSpool creates a new filament spool
func CreateFilamentSpool(c *gin.Context) {
    var filamentSpool models.FilamentSpool
    if err := c.ShouldBindJSON(&filamentSpool); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db := database.GetDB(c.Request.Context())
    if result := db.Create(&filamentSpool); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating filament spool"})
        return
    }

    c.JSON(http.StatusCreated, filamentSpool)
}

// UpdateFilamentSpool updates an existing filament spool
func UpdateFilamentSpool(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    var filamentSpool models.FilamentSpool
    db := database.GetDB(c.Request.Context())
    if result := db.First(&filamentSpool, int32(id)); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Filament spool not found"})
        return
    }

    if err := c.ShouldBindJSON(&filamentSpool); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if result := db.Save(&filamentSpool); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating filament spool"})
        return
    }

    c.JSON(http.StatusOK, filamentSpool)
}

// DeleteFilamentSpool deletes a filament spool
func DeleteFilamentSpool(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    db := database.GetDB(c.Request.Context())
    if result := db.Delete(&models.FilamentSpool{}, uint(id)); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting filament spool"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Filament spool deleted successfully"})
}