package main

import (
	"io"
	"os"
	"path"
)

type fileStorage struct {
	path string
}

// WriteFile takes the content of a file and writes it into the underlying
// storage system.
func (l fileStorage) WriteFile(filepath string, content io.Reader, contentType string) error {
	fullPath := path.Join(l.path, filepath)
	if err := os.MkdirAll(path.Dir(fullPath), 0755); err != nil {
		return err
	}

	fp, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = io.Copy(fp, content)
	return err
}

// ReadFile retrieves a file from the underlying storage, needs to return
// errFileNotFound when file is not present.
func (l fileStorage) ReadFile(filepath string) (io.ReadCloser, error) {
	fullPath := path.Join(l.path, filepath)
	if _, err := os.Stat(fullPath); err != nil {
		if os.IsNotExist(err) {
			return nil, errFileNotFound
		}
		return nil, err
	}

	return os.Open(fullPath)
}
