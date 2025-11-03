package main

import (
	"flag" // For command-line flag parsing
	"fmt"  // For printing and string formatting
	"time" // For getting the current date
)

func main() {
	// Define the name flag with default "World"
	name := flag.String("name", "Alice", "Name to greet")
	flag.Parse() // Parse the flags

	// Get current date in YYYY-MM-DD format
	today := time.Now().Format("2006-01-02")

	// Format and print the greeting
	greeting := fmt.Sprintf("Hello, %s! Today is %s.", *name, today)
	fmt.Println(greeting)
}
