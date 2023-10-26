package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Define a custom type logWriter
type logWriter struct{}

// Define the main function
func main() {
	// Make a GET request to the specified URL
	resp, err := http.Get("https://ifconfig.me/ua")
	if err != nil {
		// If there's an error, print it and exit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Create an instance of the logWriter struct
	lw := logWriter{}

	// Copy the response body to the logWriter instance
	io.Copy(lw, resp.Body)
}

// Define the Write method for the logWriter struct
func (logWriter) Write(bs []byte) (int, error) {
	// Print the byte slice as a string
	fmt.Println(string(bs))
	// Print the number of bytes written
	fmt.Println("Just wrote this many bytes:", len(bs))
	// Return the number of bytes written and nil for the error
	return len(bs), nil
}
