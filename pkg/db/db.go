package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dbName = "planetDb"

func Connection_mysql() (*gorm.DB, error) {
	// Configure the database connection (always check errors)
	dsn := "root:@(127.0.0.1:3306)/" + dbName + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return nil, err
	}
	return db, err

}
