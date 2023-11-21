package main

import "fmt"

func addNumbersInLoop(listOfNumbers[]int)(int, error){
	
	var sum int = 0

	for _, v := range listOfNumbers {
		sum += v
	}
	return sum, nil
}

func main() {

	listOfNumbers := []int{2,4,6,8,10,12,14,16,18,20}

	sum,err := addNumbersInLoop(listOfNumbers)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum)
	}
}
