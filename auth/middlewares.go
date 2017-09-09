package auth

import (
	"github.com/gin-gonic/gin"
    "errors"
)

// This middleware is responds for ServiceAuthentication.
// Fetches data from Authorization Header and returns correct Service instance.
func AuthenticationMiddleware(c *gin.Context) {
	var err error

	token := c.GetHeader("Authorization")
	service, err := ServiceAuthentication{}.AuthenticateToken(token)

	if err != nil {
		err = errors.New("API token required")
	}

	if err != nil {
		c.JSON(401, err)
		c.Abort()
		return
	}

	c.Set("service", service)
	c.Next()
}
