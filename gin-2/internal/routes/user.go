package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"payment-portal/internal/domain/user"
)

func usersRoutes(router *gin.Engine) {
	router.POST("/api/user/v1/auth", func(c *gin.Context) {

		type Person struct {
			Email string `form:"name"`
			Pass  string `form:"pass"`
		}

		var person Person

		if err := c.ShouldBind(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		//field check
		log.Println(person.Email)
		log.Println(person.Pass)

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"person":  person,
		})
	})

	router.POST("/api/internal/user/create", func(c *gin.Context) {

		var inputData user.CreateUserInput

		if err := c.ShouldBindJSON(&inputData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}


		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

}
