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

	// for _, user := range users {
	// 	db.Create(&user)
	// }

	// u := User{Username: "superman"}

	// db.Where(&u).First(&u)

	// u.LastName = "Can"

	// db.Save(&u)

	// fmt.Println(users)
}

type User struct {
	gorm.Model
	Username  string `sql:"type:VARCHAR(15);not null"`
	FirtName  string `sql:"size:100;not null"`
	LastName  string `sql:"unique;unique_index;not null;DEFAULT:'Smith'"`
	Count     int    `gorm:"AUTO_INCREMENT"`
	TempField bool   `sql:"-"`
}

var users = []User{
	User{Username: "iamsaitam", FirtName: "Ricardo", LastName: "Matias"},
	User{Username: "batman", FirtName: "Bruce", LastName: "Wayne"},
	User{Username: "superman", FirtName: "Clark", LastName: "Kent"},
}
