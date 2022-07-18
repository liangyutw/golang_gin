package models

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id       uint   `gorm:"AUTO_INCREMENT"`
	Username string `gorm:"size:100"`
	Email    string `gorm:"type:varchar(250);unique_index"`
}

func initDbByGin() (*gorm.DB, error) {
	dsn := "root:password@tcp(127.0.0.1:3306)/goblog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}

func InsertUser(array map[string]interface{}) error {

	db, err := initDbByGin()
	if err != nil {
		log.Fatal(err)
	}

	// 解析成json 字串
	// jsonString, err := json.Marshal(array)
	// if err != nil {
	// 	fmt.Printf("Error: %s", err.Error())
	// }
	// fmt.Println(string(jsonString))

	result := db.Model(&User{}).Create(array) // 通过数据的指针来创建
	fmt.Println(result.Error, result.RowsAffected)
	if result.Error != nil {
		fmt.Println("建立失敗")
	}

	return result.Error
	// return nil
}
