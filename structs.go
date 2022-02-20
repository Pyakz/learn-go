package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
	Name     string `json:"name" validate:"required" `
	Year     int    `json:"year" validate:"required"`
	Course   string `json:"course" validate:"required"`
	IDNumber string `json:"id" validate:"required"`
}

func main() {

	fmt.Println("-------- STRUCTS and VALIDATIONS --------")
	var student Student
	jsonStudent := `{"name":"Mark Vergel Banguis", "year":3, "course":"Computer Science", "id":"2017-00103"}`

	if err := json.Unmarshal([]byte(jsonStudent), &student); err != nil {
		log.Fatal(err)
	}

	fmt.Println(jsonStudent)
	fmt.Println(student)
}
