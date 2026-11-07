package controller

import (
	"Backend/config"
	"Backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NoteInput struct {
	Type  string `json:"type" binding:"required"`
	Body  string `json:"body" binding:"required"`
	Check bool   `json:"check"`
}

func GetNotes(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var notes []model.Note
	if err := config.ConnDB().Where("user_id = ?", userID).Find(&notes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve notes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"notes": notes,
	})
}

func GetByID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	noteID := c.Param("id")
	var note model.Note
	if err := config.ConnDB().Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"note": note,
	})
}

func CreateNote(c *gin.Context) {
	var input NoteInput

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note := model.Note{
		UserID: userID.(uint),
		Type:   input.Type,
		Body:   input.Body,
		Check:  input.Check,
	}

	if err := config.ConnDB().Create(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Note created successfully",
		"note":    note,
	})
}

func UpdateNote(c *gin.Context) {
	userID, _ := c.Get("user_id")
	noteID := c.Param("id")

	var note model.Note

	if err := config.ConnDB().Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	var input NoteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note.Type = input.Type
	note.Body = input.Body
	note.Check = input.Check

	if err := config.ConnDB().Save(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Note updated successfully",
		"note":    note,
	})
}

func DeleteNote(c *gin.Context) {
	userID, _ := c.Get("user_id")
	noteID := c.Param("id")

	var note model.Note

	if err := config.ConnDB().Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	if err := config.ConnDB().Delete(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Note deleted successfully",
	})
}
