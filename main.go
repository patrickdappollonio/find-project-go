package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	sourcePath   = "%s/src"
	vendorFolder = "vendor"
)

func main() {
	// Check if we have 2 args
	if len(os.Args) != 2 {
		errexit("Incorrect number of arguments passed")
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
	f, err := os.Open(gopath)

	// Check for error
	if err != nil {
		if os.IsNotExist(err) {
			errexit("$GOPATH path not found: %q", gopath)
		}

		errexit("Unknown error: %s", err.Error())
	}

	// Defer closing the file
	defer f.Close()

	// Search for the folder recursively
	pathfound, err := finduntil(f, gopath, searchingfor)

	// Check for error
	if err != nil {
		errexit("Error while searching for the folder %q: %s", searchingfor, err.Error())
	}

	// Check if we did find a path
	if pathfound == "" {
		errexit("Unable to find folder %q in $GOPATH", searchingfor)
	}

	// Print the path to screen
	fmt.Fprintln(os.Stdout, pathfound)
}

func finduntil(f *os.File, basepath, searchingfor string) (string, error) {
	// Folder is already open, read it
	// if there was a problem, we just omit this directory
	contents, _ := f.Readdir(-1)

	// Iterate over each found value
	subfolders := make([]string, 0, len(contents))

	// Try finding our own folder
	for _, c := range contents {
		// If it's not a directory or the directory is hidden, or it's a vendor folder, we don't care...
		if !c.IsDir() || filepath.HasPrefix(c.Name(), ".") || c.Name() == vendorFolder {
			continue
		}

		// Check first if this one is the one we need, so
		// we don't keep recursing
		if p := c.Name(); p == searchingfor {
			return filepath.Join(basepath, p), nil
		}

		// Append the sub-folder to the slice
		subfolders = append(subfolders, filepath.Join(basepath, c.Name()))
	}

	// Recursively check the subfolders now
	for _, sf := range subfolders {
		// Open the folder first
		sub, err := os.Open(sf)

		// Check for error, and return it if something happened
		if err != nil {
			// If the folder doesn't exist, continue
			if os.IsNotExist(err) {
				continue
			}

			return "", err
		}

		// Scan the contents of this folder too
		loc, err := finduntil(sub, sf, searchingfor)

		// Close the file
		f.Close()

		// If there was an error, return it
		if err != nil {
			return "", err
		}

		// Check if we found something
		if loc != "" {
			return loc, nil
		}
	}

	return "", nil
}

func errexit(format string, args ...interface{}) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, format)
	} else {
		fmt.Fprintf(os.Stderr, format+"\n", args...)
	}

	os.Exit(-1)
}
