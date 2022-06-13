package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	MySlice := make([]int,0,3)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		if line == "X" {
			break
		}
		numeric, err := strconv.Atoi(line)
		if err == nil {
			if len(MySlice) < cap(MySlice) {
				MySlice = append(MySlice, numeric)
			} else {
				tempMySlice := make([]int, len(MySlice)+1)
				copy(tempMySlice, MySlice)
				MySlice = tempMySlice
				MySlice[len(MySlice)-1] = numeric
			}
			sort.Ints(MySlice)
			fmt.Println(MySlice)
		}
	}
}
