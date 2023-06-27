package main

type (
  Customers struct {
    firstName string
    lastName string
    fullName string
  }
)

func getCustomers() (customers []Customers) {
  
  customers = append(customers,
  Customers{ firstName: "Abhishek", lastName: "Sharma" },
  Customers{ firstName: "Eric", lastName: "Sharma" },
  )

  return customers
}
