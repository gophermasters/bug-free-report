package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gophermasters/bug-free-report/database"
	"github.com/gophermasters/bug-free-report/database/models"
)

func ShowAllBugsReports(c *gin.Context) {
	database.Connect()
	db := database.GetDatabase()
	var p []models.Bugs
	err := db.Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find all bugs report: " + err.Error(),
		})
		return
	}
	database.CloseConn()
	c.JSON(200, p)
}

func ShowBugsReport(c *gin.Context) {
	database.Connect()
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()
	var p models.Bugs
	err = db.First(&p, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find bug report by id: " + err.Error(),
		})
		return
	}
	database.CloseConn()
	c.JSON(200, p)
}

func CreateBugsReport(c *gin.Context) {
	database.Connect()
	db := database.GetDatabase()

	var p models.Bugs

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
			"error": "cannot create bug report: " + err.Error(),
		})
		return
	}
	database.CloseConn()
	c.JSON(200, p)
}

func DeleteBugsReport(c *gin.Context) {
	database.Connect()
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Bugs{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete bug report: " + err.Error(),
		})
		return
	}
	database.CloseConn()
	c.Status(204)
}

func EditBugsReport(c *gin.Context) {
	database.Connect()
	db := database.GetDatabase()

	var p models.Bugs

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
			"error": "cannot create bug report: " + err.Error(),
		})
		return
	}
	database.CloseConn()
	c.JSON(200, p)
}