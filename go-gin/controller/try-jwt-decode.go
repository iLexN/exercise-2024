package controllers

import (
	"github.com/gin-gonic/gin"
)

func CheckJwt(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hi",
	})
}
