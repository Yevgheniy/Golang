package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main()  {
	var txtLine string
	fmt.Println("Please, type file name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fileName := scanner.Text()
	file, _ := os.Open(fileName)

	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	type myStruct struct {
		fname	string
		lname	string
	}
	MySlice := make([]myStruct,0)
	for scanner.Scan(){
		txtLine = scanner.Text()
		fullName := strings.Split(txtLine, " ")
		MySlice = append(MySlice, myStruct{fname: fullName[0], lname: fullName[1]})
	}
	file.Close()
	fmt.Println("Reading file successfully!")
	for _, item := range MySlice {
		fmt.Println(item.fname, " ", item.lname)
	}
}
