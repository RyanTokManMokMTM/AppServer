package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init(){
	dbConfig = (&database{}).Load("config/server.ini").Init()

	config := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName)
	db , err := gorm.Open(mysql.Open(config),&gorm.Config{

	})
	MusicApiDB = db //global variable
	fmt.Println(config)
	if err != nil{
		fmt.Printf("mysql connection error : %v",err)
		return
	}

	if MusicApiDB.Error != nil{
		fmt.Printf("database connection error %v",MusicApiDB.Error)
		return
	}
}