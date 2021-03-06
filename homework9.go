// import statements
package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

// this function opens and prints a binary file
func readBinaryFile(filename string) {

	// open the file
	file, err := os.Open(filename)
	// catch errors
	if err != nil {
		log.Fatal(err)
	}
	// specify to close the file later when the function is done
	defer file.Close()

	// create a new reader for the file
	reader := bufio.NewReader(file)

	// create slice of bytes
	buf := make([]byte, 256)
	// loop to read each line of the binary file, and also to catch errors
	for {
		_, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		// print each line of the file one at a time
		fmt.Printf("%s", hex.Dump(buf))
	}

}

// main function to call the function to read and print the file
func main() {

	// getting file name from input
	filename := os.Args[1]

	// calling the function to read the file
	readBinaryFile(filename)

}
