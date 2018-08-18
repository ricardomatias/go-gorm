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

	db.Model(&User{}).AddIndex("idx_first_name", "first_name")
	db.Model(&User{}).AddUniqueIndex("idx_last_name", "last_name")
}

type User struct {
	Model     gorm.Model
	Username  string
	FirstName string
	LastName  string
}

var users = []User{
	User{Username: "iamsaitam", FirstName: "Ricardo", LastName: "Matias"},
	User{Username: "batman", FirstName: "Bruce", LastName: "Wayne"},
	User{Username: "superman", FirstName: "Clark", LastName: "Kent"},
}
