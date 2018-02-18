package main

import "os"

func readdirents(osDirname string, _ []byte) ([]Dirent, error) {
	dh, err := os.Open(osDirname)
	if err != nil {
		return nil, err
	}

	fileinfos, err := dh.Readdir(0)
	if er := dh.Close(); err == nil {
		err = er
	}
	if err != nil {
		return nil, err
	}

	entries := make([]Dirent, len(fileinfos))
	for i, info := range fileinfos {
		entries[i] = &Dirent{name: info.Name(), modeType: info.Mode() & os.ModeType}
	}

	return entries, nil
}

func readdirnames(osDirname string, _ []byte) ([]string, error) {
	dh, err := os.Open(osDirname)
	if err != nil {
		return nil, err
	}

	entries, err := dh.Readdirnames(0)
	if er := dh.Close(); err == nil {
		err = er
	}
	if err != nil {
		return nil, err
	}

	return entries, nil
}
