package controllers

import (
	"github.com/Ad3bay0c/dockersAndGormTesting/db"
	"github.com/gin-gonic/gin"
)

type Application struct {
	DB *db.PostgreSql
}

func (app *Application) Index(c *gin.Context) {
	c.JSON(200, gin.H{"Message": "Working Perfectly"})
}
