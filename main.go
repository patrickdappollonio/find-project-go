package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	sourcePath   = "%s/src"
	vendorFolder = "vendor"
	logOutput    = "FIND_PROJECT_LOG"
)

func main() {
	// Create a logger
	logger := log.New(ioutil.Discard, "[find-project] ", log.Lshortfile)

	// Check if we have 2 args
	if len(os.Args) != 2 {
		errexit("Incorrect number of arguments passed")
	}

	// Enable logger if defined in the env var
	if b, _ := strconv.ParseBool(os.Getenv(logOutput)); b {
		logger.SetOutput(os.Stdout)
	}

	// Get the argument we need
	searchingfor := os.Args[1]

	// Try getting the $GOPATH
	gopath := os.Getenv("GOPATH")

	// Check if $GOPATH exists
	if gopath == "" {
		errexit("$GOPATH not set")
	}

	// Add the source path to the $GOPATH
	gopath = fmt.Sprintf(sourcePath, strings.TrimSuffix(gopath, "/"))

	// Open the directory for reading
	_, err := os.Stat(gopath)

	// Check for error
	if err != nil {
		if os.IsNotExist(err) {
			errexit("$GOPATH path not found: %q", gopath)
		}

		errexit("Unknown error: %s", err.Error())
	}

	// Search for the folder recursively
	pathfound := findpath(logger, gopath, searchingfor)

	// Check if we did find a path
	if pathfound == "" {
		errexit("Unable to find folder %q in $GOPATH", searchingfor)
	}

	// Print the path to screen
	fmt.Fprintln(os.Stdout, pathfound)
}

func errexit(format string, args ...interface{}) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, format)
	} else {
		fmt.Fprintf(os.Stderr, format+"\n", args...)
	}

	os.Exit(-1)
}
