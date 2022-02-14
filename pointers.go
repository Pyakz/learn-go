package main

import "fmt"

type Person struct {
	name    string
	age     int
	savings int
}

func main() {

	marky := Person{savings: 69000, name: "MARK", age: 21}
	sarah := Person{savings: 10000, name: "SARAH", age: 20}

	// This variables are now pointers
	// They are assign to addresses
	// These addresses points their actual values

	// pointer marky can change the values of marky
	pointerToMarky := &marky

	// pointer sarah can change the values of sarah
	pointerToSarah := &sarah

	fmt.Println("-------------------------------------------------")
	fmt.Println("-------------------------------------------------")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Printf("Value of Marky: %v\n", marky)
	fmt.Printf("Address of Marky: %p\n", &marky)
	fmt.Printf("Pointer To Marky: %v\n", pointerToMarky)
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("-------------------------------------------------")
	fmt.Println("-------------------------------------------------")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Printf("Value of Sarah: %v\n", sarah)
	fmt.Printf("Address of Sarah: %p\n", &sarah)
	fmt.Printf("Pointer To Sarah: %v\n", pointerToSarah)
	fmt.Println(" ")
	fmt.Println(" ")

	*pointerToMarky = Person{age: 21, name: "Mark Vergel B. Banguis", savings: 69}
	*pointerToSarah = Person{age: 20, name: "Sarah Mayola Mergullas", savings: 69}

	fmt.Println("-------------------------------------------------")
	fmt.Println("--------AFTER CHANGING VALUES OF POINTERS--------")
	fmt.Println("-------------------------------------------------")

	fmt.Println(" ")
	fmt.Printf("Value of Marky: %v\n", marky)
	fmt.Println(" ")
	fmt.Println("-------------------------------------------------")
	fmt.Println(" ")
	fmt.Printf("Value of Sarah: %v\n", sarah)
	fmt.Println(" ")
	fmt.Println("-------------------------------------------------")
}
