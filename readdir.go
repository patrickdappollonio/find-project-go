package main

import "os"

type dirent struct {
	name string
	mode os.FileMode
}

func (de dirent) Name() string { return de.name }
func (de dirent) IsDir() bool  { return de.mode&os.ModeDir != 0 }

func fastreaddir(osDirname string, scratchBuffer []byte) ([]*dirent, error) {
	return readdirents(osDirname, scratchBuffer)
}

func fastreaddirnames(osDirname string, scratchBuffer []byte) ([]string, error) {
	return readdirnames(osDirname, scratchBuffer)
}
