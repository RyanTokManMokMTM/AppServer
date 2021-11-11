package model

import (
	"gorm.io/gorm"
)

func init(){}

type (
	UserInfo struct {
		gorm.Model
		Username string `gorm:"username;SIZE:32;NOT NULL" json:"username" binding:"required"`
		Password string `gorm:"passoword;size:64;NOT NULL" json:"password" binding:"required""`
		Email string `gorm:"email;size:32;" json:"email"`
	}
)

// UserService TODO - CURL
type UserService interface {
	CreateUser()
	GetUserInfo(id string)
	Save()
}

func (user *UserInfo) CreateUser(){
	//TODO - Passing a data and insert to database
}

func (user *UserInfo) GetUserInfo(id string){

}

func (user* UserInfo) Save(){

}

