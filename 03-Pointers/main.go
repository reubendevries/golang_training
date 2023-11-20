package main

import "fmt"

func main() {

	var myFirstString string = "Green"

	fmt.Println("The variable myFirstString is set to", myFirstString)

	pointerConversionFunction(&myFirstString)

	fmt.Println("After passing the variable through the function pointerConversionFunction, the variable myFirstString is NOW set to", myFirstString)

}

func pointerConversionFunction(s *string) {
	var mySecondString string = "Red"

	*s = mySecondString
}
