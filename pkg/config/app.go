package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

// make sure to set your own database connection information
func Connect(){
	d, err := gorm.Open("mysql", "")
	if err != nil {
		panic("failed to connect database")
	}
	db = d

}

func GetDB() *gorm.DB {
	return db
}

