package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go1/services/jwt_service"
	"go1/services/logger"
	"net/http"
	"regexp"
)

func AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {

		bearerToken, err := getHeaderToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			logger.DefaultLogger.Error("jwt error", err)
			// the return do not stop the response.
			// so need to abort it.
			c.Abort()
			return
		}

		userInfo, err := jwt_service.Decode(bearerToken)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			logger.DefaultLogger.Error("jwt error", err)
			// the return do not stop the response.
			// so need to abort it.
			c.Abort()
			return
		}

		c.Set("userInfo", userInfo)

		c.Next()
	}
}

func getHeaderToken(c *gin.Context) (string, error) {
	authorizationHeader := c.GetHeader("Authorization")

	// Check if Authorization header is present
	if authorizationHeader == "" {
		return "", fmt.Errorf("authorization header is missing")
	}

	// Extract the Bearer token using regex
	re := regexp.MustCompile(`Bearer\s+(.*)$`)
	matches := re.FindStringSubmatch(authorizationHeader)

	if len(matches) < 2 {
		return "", fmt.Errorf("invalid Authorization header format")
	}

	bearerToken := matches[1]

	return bearerToken, nil
}
