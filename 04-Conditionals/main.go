package main

import "fmt"

func comparison() (int, error) {

	n := 0
	for n < 10 {
		n += 1
	}

	return n, nil
}

func main() {

	i, err := comparison()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The variable is set to", i)
	}
}
