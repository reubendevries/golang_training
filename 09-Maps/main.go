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

<<<<<<< HEAD
type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]User)

	me := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	myMap["me"] = me

	log.Println(myMap["me"].FirstName)
=======
func main() {
	myMap := make(map[string]int)

	myMap["First"] = 1
	myMap["Second"] = 2


	log.Println(myMap["First"])
	log.Println(myMap["Second"])
>>>>>>> refs/remotes/origin/main
}
