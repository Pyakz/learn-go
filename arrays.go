package main

import "fmt"

// func main() {

// 	fmt.Println("----- ARRAYS & SLICES -----")

// 	//array
// 	var arrays [5]int = [5]int{1, 2, 3, 4, 5}

// 	//slice [ 2 3 ]
// 	var slice []int = arrays[:]

// 	fmt.Println("Length: ", len(slice))
// 	fmt.Println("Capacity: ", cap(slice))
// 	fmt.Println("Values: ", slice)

// 	// fmt.Println("Extend: ", slice[:cap(slice)])

// }

type Student struct {
	Name string
	Year int
}

func main() {

	fmt.Println("----- SLICES -----")

	list := []Student{}

	// var list []Student = []Student{}
	// list := make([]Student, 0)

	fmt.Println(len(list))
	list = append(list, Student{Name: "Ryan Jay Enobay", Year: 22})
	list = append(list, Student{Name: "Mark Vergel B. Banguis", Year: 21})
	list = append(list, Student{Name: "Charlyn B. Banguis", Year: 20})
	list = append(list, Student{Name: "Sarah Mayola Mergullas", Year: 21})
	list = append(list, Student{Name: "AJ Banguis", Year: 24})
	list = append(list, Student{Name: "Ryan Jay Enobay", Year: 22})

	fmt.Println("----- LOOP -----")
	// Sample loop for finding a duplicate
	for i, student := range list {
		for j := i + 1; j < len(list); j++ {
			duplicate := list[j]
			if student.Name == duplicate.Name {
				fmt.Println(student.Name)
			}
		}
	}
}
