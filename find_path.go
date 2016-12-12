package main

import (
	"log"
	"os"
	"path/filepath"
	"sync"
)

func findpath(l *log.Logger, startingpoint, searchingfor string) string {
	var (
		found   = make(chan string, 1) // found holds the directory found, or empty
		folders = make(chan string)    // this channel contains a list of all the given folders to be scanned
		global  sync.WaitGroup         // global is a global list of remaining tasks, when reaching zero, it exits the function
	)

	go func() {
		// Wait and get the list of tasks for the given folders
		for current := range folders {
			go func(current string) {
				// Defer the finish of this process
				defer func() {
					l.Println("Marking folder as completed", current)
					global.Done()
				}()

				// Get all the directory folders
				dirs, err := readdirnames(current)

				// Check if there was an error, if so, log it
				// and just return
				if err != nil {
					l.Printf("*** Error while listing directory %q: %s", current, err.Error())
					return
				}

				// If there were files / folders listed, log it
				l.Println("Found", len(dirs), "files / folders names under:", current)

				// Create a local waitgroup
				var local sync.WaitGroup

				// Iterate over each one of them
				for _, dir := range dirs {
					// Add one task to the local list
					local.Add(1)
					l.Println("Adding a child task to the queue for:", dir)

					// Execute the given flow in parallel
					go func(dir, parent string) {
						// Create a path
						subpath := filepath.Join(parent, dir)

						defer func(subpath string) {
							local.Done()
							l.Println("Marking subtask as complete for:", subpath)
						}(subpath)

						// Get the stat info for this directory
						st, err := os.Stat(subpath)
						l.Println("Getting details for:", st.Name(), " -- at:", subpath)

						if err != nil {
							return
						}

						// Check if it's one of the not needed directories
						if !st.IsDir() || filepath.HasPrefix(st.Name(), ".") || st.Name() == vendorFolder {
							return
						}

						// Check if it's the name we want
						if st.Name() == searchingfor {
							found <- subpath
							return
						}

						// If not, add this subpath to the queue
						// and increase the amount of jobs in one
						global.Add(1)
						folders <- subpath
					}(dir, current)
				}

				// Wait for all local to finish before marking it as done
				local.Wait()
			}(current)
		}
	}()

	// Add the starting point and
	// add 1 to the global task list
	global.Add(1)
	folders <- startingpoint

	// This flow will wait until there's no elements
	// left in the queue, and if there's nothing left,
	// it'll send empty
	go func() {
		global.Wait()
		l.Println("Waited for global to finish. Giving up.")
		found <- ""
	}()

	// Wait until we get something or we finish looking for
	// the specific folder
	ff := <-found

	return ff
}

// readdirnames was taking from the Go std library, with the sole
// difference that this function doesn't sort the strings at the end
func readdirnames(dirname string) ([]string, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}

	names, err := f.Readdirnames(-1)
	f.Close()

	if err != nil {
		return nil, err
	}

	return names, nil
}
