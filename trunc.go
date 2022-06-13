package _main

import "fmt"

func _main() {
	var UserInput float32
	var OutputNumber int
	fmt.Printf("Please, input floating point number: ")

	fmt.Scanln(&UserInput)
	OutputNumber = int(UserInput)
	fmt.Println(OutputNumber)
}
