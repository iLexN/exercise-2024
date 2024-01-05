package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckJwt(c *gin.Context) {

	userInfo, ok := c.Get("userInfo")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token information"})
		return
	}
	c.JSON(200, gin.H{
		"userInfo": userInfo,
	})
}
