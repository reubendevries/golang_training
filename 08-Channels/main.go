package main

import "fmt"

func sum(li []int, c chan int) {

	var sum int = 0

	for _, v := range li {
		sum += v
	}

	c <- sum

}

func main() {

	myList := []int{2, 4, 6, 8, 10, 12, 16, 18, 20}
	ch := make(chan int)

	go sum(myList, ch)

	x := <-ch

	fmt.Println(x)

}
