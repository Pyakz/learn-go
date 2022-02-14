package main

import "fmt"

type person struct {
	name    string
	age     int
	address string
}

func initPerson() *person {
	newPerson := person{name: "Marky", age: 21, address: "Surigao"}
	fmt.Printf("withAddress:  %p\n", &newPerson)
	// 0xc0000a0150
	return &newPerson
}

func initPersonNoAddress() person {
	newPerson := person{name: "Pyakz", age: 21, address: "Butuan"}
	fmt.Printf("noAddress:  %p\n", &newPerson)
	// 0xc0000a01b0
	return newPerson
}

func main() {

	withAddress := initPerson()
	withoutAddress := initPersonNoAddress()

	// Address of this is equal to the address inside the function
	// 0xc0000a0150
	fmt.Printf("withAddress: %p\n", withAddress)

	// This has different address compared to the address inside the function
	// 0xc0000a0180
	fmt.Printf("noAddress: %p\n", &withoutAddress)

}
