package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Please, input string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	UserInput := scanner.Text()
	LowerUserInput := strings.ToLower(UserInput)
	if LowerUserInput[0:1] == "i" && LowerUserInput[len(LowerUserInput)-1:] == "n" && strings.Contains(LowerUserInput, "a") {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not found!")
	}
}
