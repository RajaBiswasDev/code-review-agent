package main

import (
	"fmt"
	"os"
)

const (
	maxIterations = 5
	outputFile    = "output.txt"
)

// processNumbers performs arithmetic and prints number classifications
func processNumbers(first, second int) {
	sum := first + second
	if sum > 10 {
		fmt.Println("big num")
	}

	for i := 0; i < maxIterations; i++ {
		fmt.Println("number:", i)
	}

	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		if num%2 == 0 {
			fmt.Println("even", num)
		} else {
			fmt.Println("odd", num)
		}
	}
}

// writeToFile writes a test message to the specified file
func writeToFile() error {
	f, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer f.Close()

	if _, err := f.WriteString("this is a test\n"); err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}
	return nil
}

// runTests executes the main program logic
func runTests() error {
	processNumbers(5, 10)

	if err := writeToFile(); err != nil {
		return err
	}

	fmt.Println("done")
	return nil
}
