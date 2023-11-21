package main

import "fmt"

type Person struct {
  firstName string
  lastName string
  age int
  city string
  country string
}

func main() {
	
  employee := Person{"Reuben","deVries",42,"Langley","Canada"}
  
  fmt.Printf("firstName in employee struct is set to: %s\n",employee.firstName)

  employee.firstName ="Rachel"	
  
  fmt.Printf("firstname in employee struct is now set to: %s\n",employee.firstName)
}
