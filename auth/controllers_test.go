package auth

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/slawek87/GOauth/utils"
)

func init() {
	utils.InitTestDB()
}

func TestRegisterService(t *testing.T) {
	InitMigrations()
	defer utils.CleanTestDB()

	service := Service{Name: "TestService"}
	_, err := service.RegisterService()

	if err != nil {
		t.Fail()
	}

	assert.NotEmpty(t, service.Token, "Token cannot be empty!")
}

func TestRegisterUser(t *testing.T) {
	InitMigrations()
	defer utils.CleanTestDB()

	service := Service{Name: "TestService"}
	_, err := service.RegisterService()

	if err != nil {
		t.Fail()
	}

	username := "test@test.test"
	password := "test123"

	user := User{Username: username, Password: password}
	user.RegisterUser(service)

	assert.NotEmpty(t, user.CreatedAt, "Create date cannot be empty!")
	assert.Equal(t, user.Username, username)
}

func TestResetUserPassword(t *testing.T) {
	InitMigrations()
	defer utils.CleanTestDB()

	service := Service{Name: "TestService"}
	_, err := service.RegisterService()

	if err != nil {
		t.Fail()
	}

	username := "test@test.test"
	password := "test123"

	user := User{Username: username, Password: password}
	registeredUser, _ := user.RegisterUser(service)

	registeredUser.Password = "test421"
	userChangePassword, err := registeredUser.ResetUserPassword()

	assert.NotEqual(t, registeredUser.UpdatedAt, userChangePassword.UpdatedAt, "UpdateAt cannot be the same.")
}

func TestAuthenticateUser(t *testing.T) {
	InitMigrations()
	defer utils.CleanTestDB()

	service := Service{Name: "TestService"}
	_, err := service.RegisterService()

	if err != nil {
		t.Fail()
	}

	username := "test@test.test"
	password := "test123"

	user := User{Username: username, Password: password}
	user.RegisterUser(service)

	user = User{Username: username, Password: password}
	token, err := user.AuthenticateUser()

	assert.NotEmpty(t, token,  "Token cannot be empty")
}

func TestAuthorizeUser(t *testing.T) {
	InitMigrations()
	defer utils.CleanTestDB()

	service := Service{Name: "TestService"}
	_, err := service.RegisterService()

	if err != nil {
		t.Fail()
	}

	username := "test@test.test"
	password := "test123"

	user := User{Username: username, Password: password}
	user.RegisterUser(service)

	user = User{Username: username, Password: password}
	token, err := user.AuthenticateUser()

	authorizeToken := Token{Value: token}
	assert.Equal(t, authorizeToken.AuthorizeUser(),  true,"Token value must be true")
}