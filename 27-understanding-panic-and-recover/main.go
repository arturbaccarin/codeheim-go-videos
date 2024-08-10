package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("main recovered from panic: ", r)
		}
	}()

	// Get a number from the user
	fmt.Print("enter a number: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("failed to read from standard input: ", err)
		return
	}

	input = strings.Trim(input, "\n")

	fmt.Println("processing input: ", input)
	processInput(input)

	fmt.Println("done processing input")
}

func processInput(input string) {
	defer func() {
		if r := recover(); r != nil {
			// Log the panic for debugging purposes
			log.Printf("recovered in processInput: ", r)

			// Perform any necessary cleanup
			fmt.Println("Cleaning up...")

			// Re-panic to ensure the calling function can handle it or fail gracefully
			panic(r)
		}
	}()

	if _, err := strconv.ParseInt(input, 10, 64); err != nil {
		panic(err)
	}

	fmt.Println("processed input successfully: ", input)
}

/* Best practives
* Panic only for unrecoverable errors
* Recover in top-level functions
* Clean up resources
 */
