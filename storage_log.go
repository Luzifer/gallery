package main

import (
	"bytes"
	"io"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

type logStorage struct{}

// WriteFile takes the content of a file and writes it into the underlying
// storage system.
func (l logStorage) WriteFile(filepath string, content io.Reader, contentType string) error {
	log.WithFields(log.Fields{
		"path":         filepath,
		"content_type": contentType,
	}).Info("Write file")
	return nil
}

// ReadFile retrieves a file from the underlying storage, needs to return
// errFileNotFound when file is not present.
func (l logStorage) ReadFile(filepath string) (io.ReadCloser, error) {
	log.WithFields(log.Fields{
		"path": filepath,
	}).Info("Read file")
	return ioutil.NopCloser(bytes.NewReader([]byte{})), errFileNotFound
}
