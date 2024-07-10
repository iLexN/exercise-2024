package routes

import (
	"github.com/gin-gonic/gin"
	"payment-portal/internal/container"
	"payment-portal/internal/middleware"
)

func Setup(r *gin.Engine, container *container.Container) {
	pingRoute(r)

	middlewareGroup := middleware.Middleware{
		Container: container,
	}

	usersRoutes(
		r,
		&middlewareGroup,
		container.UserRepository,
		container.JwtTokenServices,
	)

	gatewaysRoutes(
		r,
		&middlewareGroup,
		container.GatewayRepository,
	)
}
