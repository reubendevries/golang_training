package main

import (
	"golang_training/14-Packages/helpers"
	"log"
)

const numPool = 1000

func CaluclateValue(c chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	c <- randomNumber

}

func main() {

	ch := make(chan int)
	defer close(ch)

	go CaluclateValue(ch)

	num := <-ch
	log.Println(num)
}
