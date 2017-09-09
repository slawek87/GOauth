package auth

import (
	"time"
	"github.com/slawek87/GOauth/storage"
	"errors"
	"github.com/slawek87/GOauth/utils"
	"strconv"
	"strings"
)

// Reset user password.
func (user *User) ResetUserPassword() (*User, error){
    var record User

	password := utils.Hash{Password: user.Password}
	hash, err := password.GetHash()

	if err != nil {
		return user, nil
	}

	db, _ := storage.InitDB()
	db.Where(&User{Username:user.Username}).First(&record)

	record.Password = hash
	record.UpdatedAt = time.Now()

	query := db.Save(record)

	return &record, query.Error
}

// Register user in db.
func (user *User) RegisterUser(service Service) (*User, error) {
	password := utils.Hash{Password: user.Password}
	hash, err := password.GetHash()

	if err != nil {
		return user, nil
	}

	user.Service = service
	user.Password = hash
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	db, _ := storage.InitDB()
	db.NewRecord(&user)
	query := db.Create(&user)

	return user, query.Error
}


// This method compares username, password from POST request with data in db.
// The same data in db and in POST request means that user is authenticated correctly.
func (user *User) AuthenticateUser() (string, error) {
	var record User
	var tokenHistory TokenHistory
	var token Token

	hash := utils.Hash{}
	hash.Password = user.Password

	db, _ := storage.InitDB()

	if db.Where(&User{Username:user.Username}).Find(&record).Error != nil {
		return "", errors.New("Username is incorrect.")
	}

	if hash.CheckPasswordHash(record.Password) == false {
		return "", errors.New("Password is incorrect.")
	}

	tokenHistory.CreatedAt = time.Now()
	tokenHistory.UpdatedAt = time.Now()
	tokenHistory.User = record
	tokenHistory.Token = GenerateToken()

	db.NewRecord(&tokenHistory)
	query := db.Create(&tokenHistory)

	record.TokenHistoryID = tokenHistory.ID
	db.Save(&record)

	token.Key = tokenHistory.Token
	token.SetToken()

	return tokenHistory.Token, query.Error
}

// In this method we check if we have current token in db and if that token is active.
func (token *Token) AuthorizeUser() bool {
	return token.GetToken() == true
}

func (token *Token) SetToken() {
	db, _ := storage.RedisDB()
	token.Key = strings.Replace(token.Key, "-", "", -1)
	db.Set(token.Key, true, TokenExpirationTime).Err()
}


func (token *Token) GetToken() bool {
	db, _ := storage.RedisDB()
	token.Key = strings.Replace(token.Key, "-", "", -1)
	value, _ := db.Get(token.Key).Result()
	boolValue, _ := strconv.ParseBool(value)

	return boolValue
}

// method registers new company in DB.
// In this process method generates unique ServiceAuthentication TokenHistory.
// This token should be send in Authorization headers.
func (service *Service) RegisterService() (*Service, error) {
	service.Token = GenerateToken()
	service.CreatedAt = time.Now()
	service.UpdatedAt = time.Now()

	db, _ := storage.InitDB()
	db.NewRecord(&service)
	query := db.Create(&service)

	return service, query.Error
}
