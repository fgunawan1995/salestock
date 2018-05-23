package main
import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)
//Migration
func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./data.db")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&Items{}) {
		db.CreateTable(&Items{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Items{})
	}
	if !db.HasTable(&ItemsIn{}) {
		db.CreateTable(&ItemsIn{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&ItemsIn{})
	}
	if !db.HasTable(&ItemsOut{}) {
		db.CreateTable(&ItemsOut{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&ItemsOut{})
	}
	return db
}