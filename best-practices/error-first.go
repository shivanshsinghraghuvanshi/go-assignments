// Below is a snippet of code that reads a file and prints its content.
// However, it lacks proper error handling and contains unnecessary nesting.
// Refactor the code to handle errors appropriately and avoid unnecessary nesting.
package main

import (
	"fmt"
	"os"
)

func readFile() {
	file, err := os.Open("input.txt")
	if err == nil {
		// Read file content
		content := make([]byte, 1024)
		n, _ := file.Read(content)
		fmt.Println(string(content[:n]))
	} else {
		fmt.Println("Error opening file:", err)
	}
}
