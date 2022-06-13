package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main()  {
	fmt.Print("Please, input name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	Name := scanner.Text()
	fmt.Println()
	fmt.Print("Please, input address: ")
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	Address := scanner.Text()
	fmt.Println()
	myMap := map[string]string{
		"name":		Name,
		"address":	Address,
	}
	jsonString, _ := json.Marshal(myMap)
	fmt.Printf("json object in bytes: %b",jsonString)
	fmt.Println()
	fmt.Printf("json object in strings: %s",string(jsonString))
	fmt.Println()
	fmt.Println("map: ", myMap)
}

