package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", "user=gorm password=gorm dbname=gorm sslmode=disable")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.DropTable(&User{})
	db.CreateTable(&User{})

	for _, field := range db.NewScope(&User{}).Fields() {
		println(field.Name)
	}
}

type User struct {
	Model    gorm.Model `gorm:"embedded"`
	Username string
	FirtName string
	LastName string
}

var users = []User{
	User{Username: "iamsaitam", FirtName: "Ricardo", LastName: "Matias"},
	User{Username: "batman", FirtName: "Bruce", LastName: "Wayne"},
	User{Username: "superman", FirtName: "Clark", LastName: "Kent"},
}
