package middleware

import "github.com/gin-gonic/gin"

func (m *Middleware) AuthToken() gin.HandlerFunc {

	return func(c *gin.Context) {
		m.Container.Logger.Info("here is AuthToken middleware")
		c.Next()
	}
}
