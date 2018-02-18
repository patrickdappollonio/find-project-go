package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/karrick/godirwalk"
	"github.com/pkg/errors"
)

var gopath = os.Getenv("GOPATH")

type found struct {
	Path string
}

func (f *found) Error() string {
	return "No error"
}

func main() {
	if len(os.Args) != 2 {
		exit("Error: Additional parameter required: name of the folder to search for.")
	}

	searching := os.Args[1]
	gopath = filepath.Join(gopath, "src")

	err := godirwalk.Walk(gopath, &godirwalk.Options{
		Callback: func(pathName string, de *godirwalk.Dirent) error {
			if de.Name() == "vendor" || filepath.HasPrefix(de.Name(), ".") {
				return filepath.SkipDir
			}

			// fmt.Println(pathName)

			if de.Name() != searching {
				return nil
			}

			return &found{Path: pathName}
		},
	})

	if err == nil {
		exit(fmt.Sprintf("Error: Folder %q not found in $GOPATH", searching))
	}

	if data, ok := errors.Cause(err).(*found); ok {
		fmt.Fprintln(os.Stdout, data.Path)
		return
	}

	exit("Error:", err.Error(), fmt.Sprintf("Type: %T", err))
}

func exit(params ...interface{}) {
	fmt.Fprintln(os.Stderr, params...)
	os.Exit(1)
}
