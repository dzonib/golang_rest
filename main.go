package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=dzonib dbname=gopg password=123456 sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	db.Create(&User{Name: "Dzoni B", Age: 123})

}
