package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MusicApiDB *gorm.DB

const (
	DbHost string = "127.0.0.1"
	DbPort int    = 3306
	DbUser  string = "root"
	DbPassword string = "admin"
	DbTable    string = "musicDB"
)

func init(){
	config := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbTable)
	db , err := gorm.Open(mysql.Open(config),&gorm.Config{

	})
	MusicApiDB = db //global variable

	if err != nil{
		fmt.Printf("mysql connection error : %v",err)
		return
	}

	if MusicApiDB.Error != nil{
		fmt.Printf("database connection error %v",MusicApiDB.Error)
		return
	}
}
