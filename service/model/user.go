package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)
type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	PassWordDigest string
}
const (
	PassWordCost = 12 // 密码加密难度
)

// SetPassword 加密密码
func(user *User)SetPassword(password string)error{
	bytes,err:=bcrypt.GenerateFromPassword([]byte(password),PassWordCost)
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	user.PassWordDigest=string(bytes)
	return nil

}

// CheckPassword 验证密码
func(user *User)CheckPassword(password string)bool{
	err:=bcrypt.CompareHashAndPassword([]byte(user.PassWordDigest),[]byte(password))
	return  err==nil
}
