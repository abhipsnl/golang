package main

import "fmt"

func main() {

  customers := getCustomers()

  for _, customer := range customers {
    fmt.Println(customer)
  }
}
