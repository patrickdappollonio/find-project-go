package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	buf    = make([]byte, 2048)
	gopath = os.Getenv("GOPATH")
)

const usage = `Error: Missing parameter [FOLDER]

Usage: "%s [FOLDER]"
Where [FOLDER] is the name of the folder to search inside $GOPATH.
If you're using the shell alias, change %q for the
name of the alias.`

func main() {
	if len(os.Args) != 2 {
		exit(usage, os.Args[0], os.Args[0])
	}

	log.SetOutput(ioutil.Discard)
	if _, found := os.LookupEnv("FPDEBUG"); found {
		log.SetOutput(os.Stderr)
	}

	var err error
	gopath, err = filepath.Abs(filepath.Join(gopath, "src"))
	if err != nil {
		exit("Unable to get absolute path to $GOPATH, error: %s", err.Error())
	}

	loc, err := finddir(gopath, os.Args[1])
	if err != nil {
		exit("Error while traversing %q: %s", gopath, err)
	}

	if loc == "" {
		exit("Folder %q not found inside $GOPATH", os.Args[1])
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
		log.Println(dirs[i])

		if filepath.Base(dirs[i]) == name {
			return dirs[i], nil
		}

		extras, err := getalldirs(dirs[i])
		if err != nil {
			return "", err
		}

		for j := 0; j < len(extras); j++ {
			if filepath.Base(extras[j]) == name {
				log.Println(extras[j])
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
