package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func RegisterUserAPI(c *gin.Context) {
	user := User{}
	c.Bind(&user)

	service := c.MustGet("service").(Service)

	_, err := user.RegisterUser(service)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user.Password = "" // hide password

	c.JSON(http.StatusCreated, map[string]interface{}{"results": &user})
}

//
func AuthenticateUserAPI(c *gin.Context) {
	user := User{}
	c.Bind(&user)

	token, err := user.AuthenticateUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{"Token": token})
}

func AuthorizeUserAPI(c *gin.Context) {
	token := Token{Key: c.PostForm("Token")}

	c.Bind(&token)
	c.JSON(http.StatusOK, map[string]interface{}{"Authorize": token.AuthorizeUser()})
}

func ResetUserPasswordAPI(c *gin.Context) {
	user := User{}
	c.Bind(&user)

	_, err := user.ResetUserPassword()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{"Result": true})
}

// To register user you have to send POST request with Name field and its value.
func RegisterServiceAPI(c *gin.Context) {
	service := Service{}
	c.Bind(&service)

	_, err := service.RegisterService()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, &service)
}