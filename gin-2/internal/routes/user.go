package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"payment-portal/internal/domain/user"
	"payment-portal/internal/password"
)

func usersRoutes(router *gin.Engine, userRepository *user.Repository) {
	router.POST("/api/portal/user/v1/token", func(c *gin.Context) {

		type Person struct {
			Email string `json:"email" binding:"required,email"`
			Pass  string `json:"password" binding:"required"`
		}

		var person Person

		if err := c.ShouldBindJSON(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		//field check
		log.Println(person.Email)
		log.Println(person.Pass)

		loginUser, err := userRepository.GetByEmailOrName(person.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// check password
		matches, err := password.Matches(person.Pass, loginUser.Password)
		if err != nil {
			return
		}
		if !matches {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "password not match",
			})
			return
		}

		// create jwt token

		c.JSON(http.StatusOK, gin.H{
			"message":   "pong",
			"person":    person,
			"loginUser": loginUser,
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
