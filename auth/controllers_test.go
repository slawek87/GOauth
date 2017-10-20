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

	service := Service{Name: "TestingService29"}
	_, err := service.RegisterService()

	if err != nil {
		t.Fail()
	}

	assert.NotEmpty(t, service.Token, "Token cannot be empty!")
}
