package main

import "fmt"

func main() {
  customer := getCustomerData(1)
  fmt.Println(customer)
}

func getCustomerData(customerId int) (customer string) {
  var firstName = "Abhishek"
  lastName := "Sharma"

  fullname := firstName + " " + lastName

  return fullname
}
