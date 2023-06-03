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

	// Create a channel to receive the results
	ch := make(chan []models.Bugs)
	go func() {
		var p []models.Bugs
		err := db.Find(&p).Error
		if err != nil {
			ch <- nil
			return
		}
		ch <- p
	}()

	// Wait for the results from the channel
	p := <-ch

	database.CloseConn()

	if p == nil {
		c.JSON(400, gin.H{
			"error": "cannot find all bugs report",
		})
		return
	}

	c.JSON(200, p)
}

func ShowBugsReport(c *gin.Context) {
	database.Connect()
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be an integer",
		})
		return
	}

	db := database.GetDatabase()

	// Create a channel to receive the result
	ch := make(chan *models.Bugs)
	go func() {
		var p models.Bugs
		err = db.First(&p, newid).Error
		if err != nil {
			ch <- nil
			return
		}
		ch <- &p
	}()

	// Wait for the result from the channel
	p := <-ch

	database.CloseConn()

	if p == nil {
		c.JSON(400, gin.H{
			"error": "cannot find bug report by ID",
		})
		return
	}

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

	// Create a channel to receive the result
	ch := make(chan error)
	go func() {
		err = db.Create(&p).Error
		ch <- err
	}()

	// Wait for the result from the channel
	err = <-ch

	database.CloseConn()

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create bug report: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func DeleteBugsReport(c *gin.Context) {
	database.Connect()
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be an integer",
		})
		return
	}

	db := database.GetDatabase()

	// Create a channel to receive the result
	ch := make(chan error)
	go func() {
		err = db.Delete(&models.Bugs{}, newid).Error
		ch <- err
	}()

	// Wait for the result from the channel
	err = <-ch

	database.CloseConn()

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete bug report: " + err.Error(),
		})
		return
	}

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

	// Create a channel to receive the result
	ch := make(chan error)
	go func() {
		err = db.Save(&p).Error
		ch <- err
	}()

	// Wait for the result from the channel
	err = <-ch

	database.CloseConn()

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update bug report: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}
