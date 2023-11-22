package main

import (
	"log"
	"time"
)

type User struct {
	FirstName string
	MiddleName string
	LastName string
	Age int
	BirthDate time.Time
}

func(u *User) printFirstName() string {
	return u.FirstName
}

func main() {
	myMap := make(map[string]int)

	myMap["First"] = 1
	myMap["Second"] = 2


	log.Println(myMap["First"])
	log.Println(myMap["Second"])
}
