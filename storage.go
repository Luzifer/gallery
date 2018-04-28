package main

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"
)

type storage interface {
	// WriteFile takes the content of a file and writes it into the underlying
	// storage system.
	WriteFile(filepath string, content io.Reader, contentType string) error
	// ReadFile retrieves a file from the underlying storage, needs to return
	// errFileNotFound when file is not present.
	ReadFile(filepath string) (io.ReadCloser, error)
}

var errFileNotFound = errors.New("File not found")

func getStorageByURI(uri string) (storage, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	switch u.Scheme {
	case "file":
		return fileStorage{path: strings.TrimPrefix(uri, "file://")}, nil

	case "log":
		// NoOp storage: Only logs requests
		return logStorage{}, nil

	case "s3":
		return s3StorageFromURI(uri), nil

	default:
		return nil, fmt.Errorf("Storage scheme %q not defined", u.Scheme)
	}
}
