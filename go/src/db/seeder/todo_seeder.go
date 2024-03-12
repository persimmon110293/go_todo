package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Todo struct {
	Title       string
	Description string
}

func seeds(db *gorm.DB) error {

	todo1 := Todo{Title: "title 1", Description: "description 1"}

	if err := db.Create(&todo1).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	todo2 := Todo{Title: "title 2", Description: "description 2"}

	if err := db.Create(&todo2).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	todo3 := Todo{Title: "title 3", Description: "description 3"}

	if err := db.Create(&todo3).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	return nil
}

func openConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:password@tcp(mysql:3306)/practice?charset=utf8&parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}
	return db
}

func main() {
	db := openConnection()
	defer db.Close()
	if err := seeds(db); err != nil {
		fmt.Printf("%+v", err)
		return
	}
}
