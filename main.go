package main

import (
	"fmt"
	"os"
	"strings"
)

const sourcePath = "%s/src"

func main() {
	// Try getting the $GOPATH
	gopath := os.Getenv("GOPATH")

	// Check if $GOPATH exists
	if gopath == "" {
		fmt.Fprintln(os.Stderr, "$GOPATH not set")
		os.Exit(-1)
	}

	// Add the source path to the $GOPATH
	gopath = fmt.Sprintf(sourcePath, strings.TrimSuffix(gopath, "/"))

	// Check if the path exists
	if !exists(gopath) {
		fmt.Fprintf(os.Stderr, "$GOPATH path not found: %q \n", gopath)
		os.Exit(-1)
	}

	fmt.Println(gopath)

}

func exists(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}

	return true
}
