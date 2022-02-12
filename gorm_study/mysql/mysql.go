package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id      int64   `gorm:"column:id;comment:id;AUTO_INCREMENT;primary_key"`
	Name    string  `gorm:"column:name;size:50;type:varchar(100);comment:姓名"`
	City    string  `gorm:"column:city;size:10;type:varchar(50);comment:城市"`
	address string  `gorm:"column:live_address;default:beijing;comment:地址"`
	balance float32 `gorm:"column:balance;"`
}

func main() {
	// host := "127.0.0.1"
	// port := 3306
	// username := "root"
	// passport := "root"
	// db := "eastmoney"
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&User{})
	fmt.Println("db: ", db)
}
