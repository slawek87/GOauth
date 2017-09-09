package utils

import "golang.org/x/crypto/bcrypt"


type Hash struct {
	Password	string
}

func (p Hash) GetHash() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p.Password), 14)
	return string(bytes), err
}

func (p Hash) CheckPasswordHash(hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p.Password))
	return err == nil
}