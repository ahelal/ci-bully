package output

import (
	"fmt"
	"os"
)

// CheckError checks if we have an error and dies if an
func CheckError(msg string, e error) {
	if e != nil {
		fmt.Printf("%s. %s\n", msg, e.Error())
		os.Exit(1)
	}
}

// PrintError print's the msg and dies
func PrintError(msg string) {
	fmt.Printf("%s\n", msg)
	os.Exit(1)
}
