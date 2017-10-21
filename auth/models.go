package auth

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Stores all information about users.
type User struct {
	gorm.Model

	Service				Service		`gorm:"not null;unique"`

	Username 			string		`gorm:"not null;unique"`
	Password			string		`gorm:"not null;unique"`
	TokenHistoryID		uint		`gorm:"null; default: null"`
}

// Stores token history.
type TokenHistory struct {
	gorm.Model

	User   		User		`gorm:"not null"`
	Token  		string		`gorm:"not null;unique"`
}

// Stores information about Services. Services is the company.
type Service struct {
	gorm.Model

	Name		string		`gorm:"not null;unique"`
	Token		string		`gorm:"not null;unique"`

}

type Token struct {
    Value  string
	Active bool
}

var TokenExpirationTime time.Duration = 8 * time.Hour // 8hrs