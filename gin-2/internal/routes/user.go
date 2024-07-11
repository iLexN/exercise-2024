package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"payment-portal/internal/domain/user"
	"payment-portal/internal/jwt"
	"payment-portal/internal/middleware"
	"payment-portal/internal/password"
)

func usersRoutes(router *gin.Engine, mg *middleware.Middleware, userRepository *user.Repository, tokenServices *jwt.TokenServices) {
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
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if !matches {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "password not match",
			})
			return
		}

		// create jwt token
		tokenInfo, err := tokenServices.CreateToken(loginUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"access_token": tokenInfo.SignedToken,
			"expires_at":   tokenInfo.ExpireAt.UTC(),
		})
	})

	router.GET("/api/portal/user/v1/info", mg.AuthToken(), func(c *gin.Context) {
		loginUser, ok := c.Get("user")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "pong1",
			"user":    loginUser,
		})
	})

	router.POST("/api/portal/user/v1/create", mg.AuthToken(), func(c *gin.Context) {

		var inputData user.CreateUserInput

		if err := c.ShouldBindJSON(&inputData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// todo: check email is already exist?

		newUser, err := userRepository.CreateUser(&inputData)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User Created",
			"user":    newUser,
		})
	})

}
