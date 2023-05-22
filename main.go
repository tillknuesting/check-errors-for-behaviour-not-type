package main

import (
	"check-errors-for-behaviour-not-type/timeoutpkg"
	"fmt"
)

// interface for a timeout error
type timeout interface {
	Timeout() bool
}

func isTimeout(err error) bool {
	te, ok := err.(timeout)
	return ok && te.Timeout()
}

func main() {
	err := timeoutpkg.SimulateOperation()

	if err != nil {
		if isTimeout(err) {
			fmt.Println("The operation timed out, let's try again...")
			// Retry logic here
		} else {
			fmt.Println("An error occurred:", err)
		}
	}
}
