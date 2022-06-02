package database

import (
	entities "api/Entities"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbHost = "127.0.0.1"
	dbUser = "root"
	dbPass = "pass"
	dbName = "mysql"
	dbPort = "3306"
)

func New() (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName), // data source name
		DefaultStringSize:         256,                                                                                                              // default size for string fields
		DisableDatetimePrecision:  true,                                                                                                             // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                                                             // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                                                             // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                                                            // auto configure based on currently MySQL version
	}), &gorm.Config{})
	err = db.AutoMigrate(&entities.User{}, &entities.Catagory{}, &entities.Product{}, &entities.Order{}, &entities.Cart{})
	return db, err
}
