package main

import "fmt"

func animalNoise(s string) string {
	if s == "dog" {
		return "Woof"
	} else if s == "cat" {
		return "Meow"
	} else {
		return "Unknown Animal"
	}
}

func main() {

	animal := "cat"

	noise := animalNoise(animal)

	fmt.Printf("The noise a %s makes is '%s'.\n", animal, noise)

}
