package main

import (
	"fmt"
	"log"
	"os"
)

// CHANGE: Moved global variables into the functions where they are used
// and made them constants where appropriate
const (
	thresholdValue = 10
	loopCount     = 5
)

// CHANGE: Renamed function to be more descriptive of its purpose
func processNumbersAndPrint() {
	// CHANGE: Moved variables into function scope
	x := 5
	y := 10
	
	sum := x + y
	if sum > thresholdValue {
		fmt.Println("big num")
	}
	
	// CHANGE: Added comment explaining the purpose of this loop
	// Print numbers from 0 to loopCount-1
	for i := 0; i < loopCount; i++ {
		fmt.Println("number:", i)
	}

	// CHANGE: Using more idiomatic Go with range loop
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		// CHANGE: Simplified even/odd check using string format
		evenOdd := "odd"
		if num%2 == 0 {
			evenOdd = "even"
		}
		fmt.Printf("%s %d\n", evenOdd, num)
	}
}

// CHANGE: Added proper error handling and documentation
// writeToFile creates output.txt and writes a test message to it
func writeToFile() error {
	f, err := os.Create("output.txt")
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	_, err = f.WriteString("this is a test\n")
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}
	return nil
}

// CHANGE: Updated test function to handle errors
func test() {
	processNumbersAndPrint()
	if err := writeToFile(); err != nil {
		log.Printf("Error writing to file: %v", err)
	}
	fmt.Println("done")
}
