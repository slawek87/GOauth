package main

import (
	"github.com/gin-gonic/gin"
	"github.com/slawek87/GOauth/auth"
)

func main() {
	auth.InitMigrations()

	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.POST("/user/register", auth.AuthenticationMiddleware, auth.RegisterUserAPI)
		v1.POST("/user/password/reset", auth.AuthenticationMiddleware, auth.ResetUserPasswordAPI)
		v1.POST("/user/authentication", auth.AuthenticationMiddleware, auth.AuthenticateUserAPI)
		v1.POST("/user/authorization", auth.AuthenticationMiddleware, auth.AuthorizeUserAPI)

		v1.POST("/service/register", auth.RegisterServiceAPI)
	}

	r.Run(":8090")
}

