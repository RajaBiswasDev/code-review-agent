package main

import (
	"fmt"
	"os"
)

var x = 5
var y = 10

func doingThings() {
	a := x + y
	if a > 10 {
		fmt.Println("big num")
	}
	for i := 0; i < 5; i++ {
		fmt.Println("number:", i)
	}

	stuff := []int{1, 2, 3, 4, 5}
	for j := 0; j < len(stuff); j++ {
		if stuff[j]%2 == 0 {
			fmt.Println("even", stuff[j])
		} else {
			fmt.Println("odd", stuff[j])
		}
	}
}

func writeToFile() {
	f, _ := os.Create("output.txt")
	f.WriteString("this is a test\n")
	defer f.Close()
}

func main() {
	doingThings()
	writeToFile()
	fmt.Println("done")
}
