package routes

import (
	"github.com/gin-gonic/gin"
	"payment-portal/internal/container"
)

func Setup(r *gin.Engine, container *container.Container) {
	pingRoute(r)

	usersRoutes(r, container.UserRepository)
}
