package main

import (
	"fmt"
	"time"

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

	db.DropTable(&Calendar{})
	db.CreateTable(&Calendar{})

	db.DropTable(&Appointment{})
	db.CreateTable(&Appointment{})

	users := []User{
		{Username: "robocop"},
		{Username: "David Bowie"},
		{Username: "super-derp"},
	}

	for i := range users {
		db.Save(&users[i])
	}

	db.Debug().Save(&User{
		Username: "batman",
		Calendar: Calendar{
			Name: "Improbable Events",
			Appointments: []Appointment{
				{
					Subject:     "birthday",
					Description: "foo",
					Attendees:   users,
				},
				{
					Subject:     "death",
					Description: "foo",
					Attendees:   users,
				},
			},
		},
	})

	u := User{}
	c := Calendar{}

	db.First(&u).Related(&c, "calendar")

	fmt.Println(c)
}

type User struct {
	gorm.Model
	Username  string
	FirstName string
	LastName  string
	Calendar  Calendar
}

type Calendar struct {
	gorm.Model
	Name         string
	UserID       uint
	Appointments []Appointment `gorm:"polymorphic:owner"`
}

type Appointment struct {
	gorm.Model
	Subject     string
	Description string
	StartTime   time.Time
	Length      uint
	CalendarID  uint
	OwnerID     uint
	OwnerType   string
	Attendees   []User `gorm:"many2many:appointment_user"`
}

type TaskList struct {
	gorm.Model
	Appointments []Appointment `gorm:"polymorphic:owner"`
}
