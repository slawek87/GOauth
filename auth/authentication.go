package auth

import (
	"encoding/base64"
	"github.com/satori/go.uuid"
	"errors"
	"strings"
	"github.com/slawek87/GOauth/storage"
)


type ServiceAuthentication struct {}


// Use this method to authenticate service token.
// Method is used in AuthenticationMiddleware.
func (authentication *ServiceAuthentication) decodeToken(token string) (string, string, error) {
	splitToken := strings.Split(token, " ")

	if len(splitToken) == 0 {
		return "", "", errors.New("Empty token.")
	}

	decodeToken, err := base64.StdEncoding.DecodeString(splitToken[1])

	if err != nil {
		return "", "", err
	}

	splitDecodeToken := strings.Split(string(decodeToken), ":")

	return splitDecodeToken[0], splitDecodeToken[1], err
}

func (authentication ServiceAuthentication) AuthenticateToken(token string) (Service, error) {
	name, decodeToken, err := authentication.decodeToken(token)
	service := Service{}

	if err != nil {
		return service, err
	}

	var record Service

	service.Name = name
	service.Token = decodeToken

	db, _ := storage.InitDB()

	db.Where(&Service{Name: service.Name, Token: service.Token}).First(&record)

	if err != nil || &record == nil {
		return service, errors.New("TokenHistory isn't correct.")
	}

	return record, nil
}


func GenerateToken() string {
	return uuid.NewV4().String()
}

