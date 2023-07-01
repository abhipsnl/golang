package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {

	devops := Employee{
		firstName: "Abhishek",
		lastName:  "Sharma",
		id:        101,
	}

	sre := Employee{
		"Hina",
		"Sharma",
		102,
	}

	var emp3 Employee
	emp3.firstName = "Eric"
	emp3.lastName = "Sharma"
	emp3.id = 103

	fmt.Println(devops)
	fmt.Println(sre)
	fmt.Println(emp3)

}
