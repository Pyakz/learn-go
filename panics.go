package main

import "log"

func main() {

	// names := []string{
	// 	"lobster",
	// 	"sea urchin",
	// 	"sea cucumber",
	// }
	// fmt.Println("My favorite sea creature is:", names[len(names)])
	// fmt.Println("Defer")

	// fmt.Println("1. Starting")

	// // even though defer are before the program ends
	// // Defer will be called before the panic takes place

	// defer fmt.Println("2. Defer Run")

	// // panics exits the function or ends it
	// panic("Cannot continue the program")

	// s := "100"
	defer log.Print("3. Afteeeeeeer Defeeeeeeeer......")

	defer func() {
		log.Print("1. Defeeeeeeeeeeeeeeeeeeeeeeeeeeeeeer......")

		if r := recover(); r != nil {
			// Panic will not happen
			log.Print("2. Recoooooooooooooooooooooooover......")
		}
	}()

	panic("Panic will end the program")
}
