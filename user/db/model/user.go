package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PassWordDigest string `gorm:"not null"`
}

const (
	PassWordCost = 12 //密码加密难度
)

// SetPassWord 加密密码
func (user *User) SetPassWord(pwd string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), PassWordCost)
	if err != nil {
		user.PassWordDigest = string(bytes)
	}
	return err
}

// CheckPassWord  检验密码
func (user *User) CheckPassWord(pwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.PassWordDigest), []byte(pwd)) == nil
}
