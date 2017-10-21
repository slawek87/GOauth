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

	assert.NotEmpty(t, service.Token, "Token cannot be empty!")

	username := "test@test.test"
	password := "test123"

	user := User{Username: username, Password: password}
	user.RegisterUser(service)

	assert.NotEmpty(t, user.CreatedAt, "Create date cannot be empty!")
	assert.Equal(t, user.Username, username)
}
