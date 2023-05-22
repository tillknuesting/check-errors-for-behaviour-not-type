package main

import (
	"fmt"
	"net"
	"os"
)

// We create an interface for a temporary error
type temporary interface {
	Temporary() bool
}

func isTemporary(err error) bool {
	te, ok := err.(temporary)
	return ok && te.Temporary()
}

// Use the standard library's net.Dial function as an example
func connectToServer() {
	conn, err := net.Dial("tcp", "example.com:http")
	if err != nil {
		fmt.Println("Error occurred:", err)

		// Here we inspect if the error is temporary
		if isTemporary(err) {
			fmt.Println("Don't worry, this error is temporary. Let's retry...")

			// Then we can add a retry logic here
		} else {
			fmt.Println("This error is serious. Let's stop the operation.")
			os.Exit(1)
		}
	}

	fmt.Println("Connection established:", conn)
}

func main() {
	connectToServer()
}
