package main

import (
	"fmt"
	"strings"
)

// func handle(err error) (string, error) {
// 	return "",  err
// }


type InvalidError  struct {
	Message string
	Code int
}

func (err InvalidError) Error() string {
	return fmt.Sprintf("Details: %s\nResponse Code: %d",err.Message,err.Code)
}


func capitalize(name string) (string,  error) {

	if name == "" {
		return name, InvalidError{Message: "Cannot format empty string", Code: 404}
	}

	return strings.ToTitle(name),  nil
}

func Login(url string, status bool)(string,error) {
	if status {
		return "Success", nil
	}

	return "Failed",InvalidError{Message: "Error, something happened", Code: 401}
}

func main() {

	status, err := Login("sample.com/api", false)
	// status, err := Login("sample.com/api", true)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Status: %s\n", status)



}


// Both fmt.Errorf() & errors.New()
// Both of these functions allow you to specify 
// a custom error message that you can later present to your users.