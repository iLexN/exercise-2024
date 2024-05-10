package routes

import "github.com/gin-gonic/gin"

func Setup(r *gin.Engine) {
	pingRoute(r)

	usersRoutes(r)
}
