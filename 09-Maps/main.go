package main

import "log"

func main() {
	myMap := make(map[string]int)

	myMap["First"] = 1
	myMap["Second"] = 2

	log.Println(myMap["First"], myMap["Second"])
}
