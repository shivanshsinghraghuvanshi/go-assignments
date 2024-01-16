// The code attempts to write data to a file but lacks proper error handling. Additionally, it does not remove partial data in case of a write failure.
// Refactor the code to handle errors effectively and remove any partial data written to the file.
package main

import (
	"fmt"
	"os"
)

func writeToFile(data string) {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	// TODO: Write data to the file

	// TODO: Properly handle errors and remove partial data if necessary
}

func main() {
	data := "This is some example data."
	writeToFile(data)
}
