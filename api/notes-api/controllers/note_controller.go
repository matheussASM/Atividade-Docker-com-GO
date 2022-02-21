package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/matheussASM/pre-atividade-02/database"
	"github.com/matheussASM/pre-atividade-02/models"
)

func ShowAllNotes(c *gin.Context) {
	db := database.GetDatabase()
	var p []models.Note
	err := db.Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func ShowNote(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()
	var p models.Note
	err = db.First(&p, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func CreateNote(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Note

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create Note: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func DeleteNote(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Note{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete Note: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func EditNote(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Note

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Save(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create Note: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}
