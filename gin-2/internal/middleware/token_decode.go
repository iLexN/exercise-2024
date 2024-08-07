package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
)

var errAuthHeaderMissing = errors.New("authorization header is missing")
var errBearerMissing = errors.New("bearer token missing")

func (m *Middleware) AuthToken() gin.HandlerFunc {

	return func(c *gin.Context) {
		m.Container.Logger.Info("here is AuthToken middleware")

		bearer, err := getBearerToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
		}

		claims, err := m.Container.JwtTokenServices.DecodeToken(bearer)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
		}

		id, err := strconv.ParseUint(claims.ID, 10, 32)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		user, err := m.Container.UserRepository.GetById(uint(id))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("user", user)

		c.Next()
	}
}

func getBearerToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		return "", errAuthHeaderMissing
	}

	re := regexp.MustCompile(`Bearer\s+(.*)$`)
	matches := re.FindStringSubmatch(authHeader)

	if len(matches) < 2 {
		return "", errBearerMissing
	}

	bearerToken := matches[1]

	return bearerToken, nil
}
