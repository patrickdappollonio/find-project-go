package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	buf    = make([]byte, 2048)
	gopath = os.Getenv("GOPATH")
)

func main() {
	if len(os.Args) != 2 {
		exit("Error: Additional parameter required: name of the folder to search for.")
	}

	loc, err := finddir(filepath.Join(gopath, "src"), os.Args[1])
	if err != nil {
		exit("Error while traversing $GOPATH at %q: %s", gopath, err)
	}

	fmt.Fprintln(os.Stdout, loc)
}

func exit(format string, params ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", params...)
	os.Exit(1)
}

func finddir(p, name string) (string, error) {
	var dirs []string
	var err error

	dirs, err = getalldirs(p)
	if err != nil {
		return "", err
	}

	for i := 0; i < len(dirs); i++ {
		if filepath.Base(dirs[i]) == name {
			return dirs[i], nil
		}

		extras, err := getalldirs(dirs[i])
		if err != nil {
			return "", err
		}

		for j := 0; j < len(extras); j++ {
			if filepath.Base(extras[j]) == name {
				return extras[j], nil
			}

			dirs = append(dirs, extras[j])
		}
	}

	return "", nil
}

func getalldirs(p string) ([]string, error) {
	ls, err := fastreaddir(p, buf)
	if err != nil {
		return nil, err
	}

	var all []string
	for i := 0; i < len(ls); i++ {
		if !ls[i].IsDir() {
			continue
		}

		if ls[i].Name() == "vendor" || filepath.HasPrefix(ls[i].Name(), ".") {
			continue
		}

		all = append(all, filepath.Join(p, ls[i].Name()))
	}

	return all, nil
}
